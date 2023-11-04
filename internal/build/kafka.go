package build

import (
	"github.com/IBM/sarama"
	"github.com/pkg/errors"
)

func (b *Builder) kafkaSyncProducer() (sarama.SyncProducer, error) {
	config, err := b.kafkaConfig()
	if err != nil {
		return nil, err
	}

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	config.ClientID = b.config.App.Name

	producer, err := sarama.NewSyncProducer(b.config.Kafka.Brokers, config)
	if err != nil {
		return nil, errors.Wrap(err, "build kafka producer")
	}

	return producer, nil
}

func (b *Builder) kafkaConfig() (*sarama.Config, error) {
	config := sarama.NewConfig()

	version, err := sarama.ParseKafkaVersion(b.config.Kafka.Version)
	if err != nil {
		return nil, errors.Wrapf(err, "parsing kafka version: %q", b.config.Kafka.Version)
	}

	config.Version = version

	config.Net.SASL.Enable = b.config.Kafka.SASL()
	config.Net.SASL.User = b.config.Kafka.User
	config.Net.SASL.Password = b.config.Kafka.Password
	config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	config.Net.SASL.Handshake = true

	return config, nil
}
