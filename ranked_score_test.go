package midjourney

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRankedScores_MarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		rs   RankedScores
		want string
	}{
		{
			name: "empty",
			rs:   RankedScores{},
			want: `""`,
		},
		{
			name: "one score",
			rs:   RankedScores{Mehd},
			want: `"2"`,
		},
		{
			name: "multiple scores",
			rs:   RankedScores{Unranked, Loved},
			want: `"0,5"`,
		},
		{
			name: "all scores",
			rs:   RankedScores{Unranked, Mehd, Liked, Loved},
			want: `"0,2,4,5"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.rs)
			require.NoError(t, err)

			assert.Equal(t, tt.want, string(got))
		})
	}
}

func TestRankedScores_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		json string
		want RankedScores
	}{
		{
			name: "empty",
			json: `""`,
			want: nil,
		},
		{
			name: "one score",
			json: `"2"`,
			want: RankedScores{Mehd},
		},
		{
			name: "multiple scores",
			json: `"0,5"`,
			want: RankedScores{Unranked, Loved},
		},
		{
			name: "all scores",
			json: `"0,2,4,5"`,
			want: RankedScores{Unranked, Mehd, Liked, Loved},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got RankedScores
			err := json.Unmarshal([]byte(tt.json), &got)
			require.NoError(t, err)

			assert.Equal(t, tt.want, got)
		})
	}
}
