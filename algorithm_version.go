package midjourney

import (
	"encoding/json"
	"strconv"
)

type AlgorithmVersion string

func (av *AlgorithmVersion) MarshalJSON() ([]byte, error) {
	n, err := strconv.Atoi(string(*av))
	if err != nil {
		return json.Marshal(string(*av))
	}

	return json.Marshal(n)
}

func (av *AlgorithmVersion) UnmarshalJSON(b []byte) error {
	var n int
	err := json.Unmarshal(b, &n)
	if err == nil {
		*av = AlgorithmVersion(strconv.Itoa(n))

		return nil
	}

	var s string
	err = json.Unmarshal(b, &s)
	if err == nil {
		*av = AlgorithmVersion(s)

		return nil
	}

	return err
}
