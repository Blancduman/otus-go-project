package slot

import "errors"

var (
	ErrNotFound             = errors.New("slot not found")
	ErrInvalidSlotIDCreated = errors.New("slot created with invalid id")
)
