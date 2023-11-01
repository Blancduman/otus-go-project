package banner

import "github.com/pkg/errors"

var (
	ErrNotFound               = errors.New("banner not found")
	ErrInvalidBannerIDCreated = errors.New("banner created with invalid id")
)
