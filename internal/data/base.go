package data

import (
	"io"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-json"
)

type baseReader struct{}

func (b baseReader) ToReader() (io.Reader, error) {
	bytes, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(bytes)), nil
}

type baseValidator struct{}

func (b baseValidator) Validate() error {
	return validator.New().Struct(b)
}
