package config

type Kafka struct {
	Version  string   `envconfig:"KAFKA_VERSION" default:"1.0.0"`
	User     string   `envconfig:"KAFKA_USER"`
	Password string   `envconfig:"KAFKA_PASSWORD"`
	Brokers  []string `envconfig:"KAFKA_BROKERS"`
	Topics   KafkaTopics
}

type KafkaTopics struct {
	Stat string `envconfig:"KAFKA_TOPIC_STAT"`
}

func (k Kafka) SASL() bool {
	return k.User != "" && k.Password != ""
}
