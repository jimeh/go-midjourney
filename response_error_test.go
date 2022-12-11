package midjourney

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseError_Is(t *testing.T) {
	tests := []struct {
		name string
		is   error
		want bool
	}{
		{
			name: "Err",
			is:   Err,
			want: true,
		},
		{
			name: "ErrResponse",
			is:   ErrResponse,
			want: true,
		},
		{
			name: "ErrResponse",
			is:   ErrResponse,
			want: true,
		},
		{
			name: "ErrInvalidAPIURL",
			is:   ErrInvalidAPIURL,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respErr := &ResponseError{Message: "foo"}

			got := errors.Is(respErr, tt.is)

			assert.Equal(t, tt.want, got)
		})
	}
}
