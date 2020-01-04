# robocar-objects-detection

Microservice that detect objects on camera frames
     
## Usage
```
rc-objects-detection <OPTIONS>

  -mqtt-broker string
        Broker Uri, use MQTT_BROKER env if arg not set (default "tcp://127.0.0.1:1883")
  -mqtt-client-id string
        Mqtt client id, use MQTT_CLIENT_ID env if args not set (default "robocar-objects-detection")
  -mqtt-password string
        Broker Password, MQTT_PASSWORD env if args not set
  -mqtt-qos int
        Qos to pusblish message, use MQTT_QOS env if arg not set
  -mqtt-retain
        Retain mqtt message, if not set, true if MQTT_RETAIN env variable is set
  -mqtt-topic-frame string
        Mqtt topic that contains camera frame, use MQTT_TOPIC_CAMERA if args not set
  -mqtt-topic-objects string
        Mqtt topic to publish discovered objects, use MQTT_TOPIC_OBJECT if args not set
  -mqtt-username string
        Broker Username, use MQTT_USERNAME env if arg not set
  -tf-model-config-path string
        Tensorflow config model path, use TF_MODEL_CONFIG_PATH if args not set
  -tf-model-path string
        Tensorflow model path, use TF_MODEL_PATH if args not set
```

## Docker build

```bash
export DOCKER_CLI_EXPERIMENTAL=enabled
docker buildx build . --platform linux/amd64,linux/arm/7,linux/arm64 -t cyrilix/robocar-throttle
```
