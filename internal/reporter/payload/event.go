package payload

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Message struct {
	Type    Type            `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type Type string

const (
	TypeClick Type = "click"
	TypeShown Type = "shown"
)

type Payload interface {
	Type() Type
}

func Encode(p Payload) ([]byte, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, errors.Wrap(err, "encode payload")
	}

	b, err := json.Marshal(Message{
		Type:    p.Type(),
		Payload: data,
	})
	if err != nil {
		return nil, errors.Wrap(err, "encode message")
	}

	return b, nil
}
