package main

type ExampleCreateConnector struct {
	Name   string                       `json:"name"`
	Config ExampleCreateConnectorConfig `json:"config"`
}

type ExampleCreateConnectorConfig struct {
	ConnectorClass                   string `json:"connector.class"`
	TasksMax                         int    `json:"tasks.max"`
	MqttServerUri                    string `json:"mqtt.server.uri"`
	MqttUsername                     string `json:"mqtt.username"`
	MqttPassword                     string `json:"mqtt.password"`
	MqttTopics                       string `json:"mqtt.topics"`
	KafkaTopic                       string `json:"kafka.topic"`
	ValueConverter                   string `json:"value.converter"`
	ConfluentTopicBootstrapServers   string `json:"confluent.topic.bootstrap.servers"`
	ConfluentTopicReplicationFactor  int    `json:"confluent.topic.replication.factor"`
	ConnectMqttConverterThrowOnError string `json:"connect.mqtt.converter.throw.on.error"`
}
