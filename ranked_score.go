package midjourney

import (
	"encoding/json"
	"strconv"
	"strings"
)

type RankedScores []RankedScore

// URIParam returns a string representation of the RankedScores suitable for URI
// parameters.
func (rs RankedScores) URIParam() string {
	vals := make([]string, 0, len(rs))

	for _, v := range rs {
		vals = append(vals, strconv.Itoa(int(v)))
	}

	return strings.Join(vals, ",")
}

func (rs RankedScores) MarshalJSON() ([]byte, error) {
	return json.Marshal(rs.URIParam())
}

func (rs *RankedScores) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	scores := strings.Split(s, ",")

	for _, score := range scores {
		if score == "" {
			continue
		}

		val, err := strconv.Atoi(score)
		if err != nil {
			return err
		}

		*rs = append(*rs, RankedScore(val))
	}

	return nil
}

type RankedScore int

const (
	Unranked RankedScore = 0
	Mehd     RankedScore = 2
	Liked    RankedScore = 4
	Loved    RankedScore = 5
)

func (rs RankedScore) String() string {
	switch rs {
	case Mehd:
		return "meh"
	case Liked:
		return "liked"
	case Loved:
		return "loved"
	case Unranked:
		return "unranked"
	default:
		return "unknown"
	}
}

func (rs RankedScore) URIParam() string {
	return strconv.Itoa(int(rs))
}
