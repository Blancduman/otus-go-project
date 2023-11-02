package reporter

import (
	"context"
	"time"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/Blancduman/banners-rotation/internal/reporter/payload"
)

type Producer struct {
	SaramaSyncProducer sarama.SyncProducer
	Topic              string
}

func (p *Producer) Produce(_ context.Context, pl payload.Payload) error {
	key := uuid.New()

	value, err := payload.Encode(pl)
	if err != nil {
		return errors.Wrapf(err, "encode %s payload", pl.Type())
	}

	// задать заголовки, метаданные и т.д. можно и нужно с помощью мидлвар для sarama.SyncProducer
	// не юзай watermill и т.п.
	_, _, err = p.SaramaSyncProducer.SendMessage(&sarama.ProducerMessage{
		Topic:     p.Topic,
		Key:       sarama.ByteEncoder(key[:]),
		Value:     sarama.ByteEncoder(value),
		Headers:   nil,
		Metadata:  nil,
		Offset:    0,
		Partition: 0,
		Timestamp: time.Time{},
	})

	return errors.Wrapf(err, "send %s change message", pl.Type())
}
