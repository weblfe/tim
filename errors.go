package tim

import "errors"

var (
	ErrInvalidPriKey = errors.New("Invalid private key")
	ErrInvalidVer    = errors.New("Unknown version")
)
