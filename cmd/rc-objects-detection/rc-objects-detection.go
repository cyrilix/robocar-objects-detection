package main

import (
	"flag"
	"github.com/cyrilix/robocar-base/cli"
	"github.com/cyrilix/robocar-objects-detection/objects"
	"github.com/cyrilix/robocar-objects-detection/part"
	"go.uber.org/zap"
	"log"
	"os"
)

const (
	DefaultClientId = "robocar-objects-detection"
)

func main() {
	var mqttBroker, username, password, clientId string
	var disparityTopic, objectsTopic, objectsCleanTopic string
	var imgWidth, imgHeight int
	var objectSizeThreshold float64

	mqttQos := cli.InitIntFlag("MQTT_QOS", 0)
	_, mqttRetain := os.LookupEnv("MQTT_RETAIN")

	cli.InitMqttFlags(DefaultClientId, &mqttBroker, &username, &password, &clientId, &mqttQos, &mqttRetain)

	flag.StringVar(&disparityTopic, "mqtt-topic-frame", os.Getenv("MQTT_TOPIC_DISPARITY"), "Mqtt topic that contains disparity frame, use MQTT_TOPIC_DISPARITY if args not set")
	flag.StringVar(&objectsTopic, "mqtt-topic-objects", os.Getenv("MQTT_TOPIC_OBJECTS"), "Mqtt topic to publish discovered objects, use MQTT_TOPIC_OBJECT if args not set")
	flag.StringVar(&objectsCleanTopic, "mqtt-topic-objects-clean", os.Getenv("MQTT_TOPIC_OBJECTS_CLEAN"), "Mqtt topic to publish filtered objects, use MQTT_TOPIC_OBJECT_CLEAN if args not set")

	flag.IntVar(&imgWidth, "image-width", 160, "Video pixels width")
	flag.IntVar(&imgHeight, "image-height", 128, "Video pixels height")
	flag.Float64Var(&objectSizeThreshold, "object-size-threshold", 0.75, "Max object size in percent of image to filter")

	logLevel := zap.LevelFlag("log", zap.InfoLevel, "log level")
	flag.Parse()

	if len(os.Args) <= 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(*logLevel)
	lgr, err := config.Build()
	if err != nil {
		log.Fatalf("unable to init logger: %v", err)
	}
	defer func() {
		if err := lgr.Sync(); err != nil {
			log.Printf("unable to Sync logger: %v\n", err)
		}
	}()
	zap.ReplaceGlobals(lgr)

	client, err := cli.Connect(mqttBroker, username, password, clientId)
	if err != nil {
		log.Fatalf("unable to connect to mqtt bus: %v", err)
	}
	defer client.Disconnect(50)

	processor := objects.NewFilter(imgWidth, imgHeight, objectSizeThreshold)
	p := part.New(client, disparityTopic, objectsTopic, objectsCleanTopic, processor)
	defer p.Stop()

	cli.HandleExit(p)

	err = p.Start()
	if err != nil {
		log.Fatalf("unable to start service: %v", err)
	}
}
