package config

type Kafka struct {
	Version        string   `envconfig:"KAFKA_VERSION" default:"1.0.0"`
	User           string   `envconfig:"KAFKA_USER"`
	Password       string   `envconfig:"KAFKA_PASSWORD"`
	Brokers        []string `envconfig:"KAFKA_BROKERS"`
	TopicUpdating  string   `envconfig:"KAFKA_TOPIC_UPDATING"`
	TopicDeleting  string   `envconfig:"KAFKA_TOPIC_DELETING"`
	ConsumerGroup  string   `envconfig:"KAFKA_CONSUMER_GROUP"`
	TopicProducing string   `envconfig:"KAFKA_TOPIC_PRODUCING"`
}

func (k Kafka) SASL() bool {
	return k.User != "" && k.Password != ""
}
