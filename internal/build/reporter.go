package build

import (
	"github.com/Blancduman/banners-rotation/internal/reporter"
)

func (b *Builder) reporterProducer() (*reporter.Producer, error) {
	saramaProducer, err := b.kafkaSyncProducer()
	if err != nil {
		return nil, err
	}

	return &reporter.Producer{
		SaramaSyncProducer: saramaProducer,
		Topic:              b.config.Kafka.Topics.Stat,
	}, nil
}
