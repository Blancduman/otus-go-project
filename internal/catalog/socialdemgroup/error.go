package socialdemgroup

import "github.com/pkg/errors"

var (
	ErrNotFound                       = errors.New("social dem group not found")
	ErrInvalidSocialDemGroupIDCreated = errors.New("social dem group created with invalid id")
)
