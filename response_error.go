package midjourney

import (
	"fmt"
)

type ResponseError struct {
	Message string `json:"error,omitempty"`
	message string
}

func (re *ResponseError) Error() string {
	if re.message != "" {
		return re.message
	}
	re.message = fmt.Errorf("%w: %s", ErrResponse, re.Message).Error()

	return re.message
}

func (re *ResponseError) Unwrap() error {
	return ErrResponse
}
