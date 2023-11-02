package config

type Kafka struct {
	Version        string   `envconfig:"KAFKA_VERSION" default:"1.0.0"`
	User           string   `envconfig:"KAFKA_USER"`
	Password       string   `envconfig:"KAFKA_PASSWORD"`
	Brokers        []string `envconfig:"KAFKA_BROKERS"`
	Topics         KafkaTopics
	SchemaRegistry KafkaSchemaRegistry
}

type KafkaTopics struct {
	Stat string `envconfig:"KAFKA_TOPIC_STAT"`
}

type KafkaSchemaRegistry struct {
	URL string `envconfig:"KAFKA_SCHEMA_REGISTRY_URL"`
}

func (k Kafka) SASL() bool {
	return k.User != "" && k.Password != ""
}
