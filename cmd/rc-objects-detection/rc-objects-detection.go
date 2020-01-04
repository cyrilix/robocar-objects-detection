package main

import (
	"flag"
	"github.com/cyrilix/robocar-base/cli"
	"github.com/cyrilix/robocar-objects-detection/part"
	"log"
	"os"
)

const (
	DefaultClientId = "robocar-objects-detection"
)

func main() {
	var mqttBroker, username, password, clientId string
	var cameraTopic, objectsTopic string
	var modelPath, modelConfigPath string

	mqttQos := cli.InitIntFlag("MQTT_QOS", 0)
	_, mqttRetain := os.LookupEnv("MQTT_RETAIN")

	cli.InitMqttFlags(DefaultClientId, &mqttBroker, &username, &password, &clientId, &mqttQos, &mqttRetain)

	flag.StringVar(&cameraTopic, "mqtt-topic-frame", os.Getenv("MQTT_TOPIC_CAMERA"), "Mqtt topic that contains camera frame, use MQTT_TOPIC_CAMERA if args not set")
	flag.StringVar(&objectsTopic, "mqtt-topic-objects", os.Getenv("MQTT_TOPIC_OBJECTS"), "Mqtt topic to publish discovered objects, use MQTT_TOPIC_OBJECT if args not set")
	flag.StringVar(&modelPath, "tf-model-path", os.Getenv("TF_MODEL_PATH"), "Tensorflow model path, use TF_MODEL_PATH if args not set")
	flag.StringVar(&modelConfigPath, "tf-model-config-path", os.Getenv("TF_MODEL_CONFIG_PATH"), "Tensorflow config model path, use TF_MODEL_CONFIG_PATH if args not set")

	flag.Parse()
	if len(os.Args) <= 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	client, err := cli.Connect(mqttBroker, username, password, clientId)
	if err != nil {
		log.Fatalf("unable to connect to mqtt bus: %v", err)
	}
	defer client.Disconnect(50)

	detector := part.NewCarDetector(modelPath, modelConfigPath)
	p := part.New(client, cameraTopic, objectsTopic, detector)
	defer p.Stop()

	cli.HandleExit(p)

	err = p.Start()
	if err != nil {
		log.Fatalf("unable to start service: %v", err)
	}
}
