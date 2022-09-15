package midjourney

import (
	"fmt"
	"strings"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05.999999"

type Time struct {
	time.Time
}

func (ct *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}

		return
	}

	ct.Time, err = time.Parse(TimeFormat, s)

	return
}

func (ct *Time) MarshalJSON() ([]byte, error) {
	if ct.Time.IsZero() {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(TimeFormat))), nil
}
