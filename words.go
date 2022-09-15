package midjourney

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
)

type Word struct {
	Word    string
	ImageID string
}

func (w *Word) ImageURL() string {
	return fmt.Sprintf("https://i.mj.run/%s/0_0.png", w.ImageID)
}

type WordsQuery struct {
	Query  string
	Amount int
	Page   int
	Seed   int

	// RandomSeed will send a random Seed value between 0 and 9999.
	RandomSeed bool
}

func (rjq *WordsQuery) Values() url.Values {
	v := url.Values{}
	if rjq.Query != "" {
		v.Set("query", rjq.Query)
	}
	if rjq.Amount != 0 {
		v.Set("amount", strconv.Itoa(rjq.Amount))
	}
	v.Set("page", strconv.Itoa(rjq.Page))
	if rjq.RandomSeed {
		v.Set("seed", strconv.Itoa(randInt(9999)))
	} else if rjq.Seed != 0 {
		v.Set("seed", strconv.Itoa(rjq.Seed))
	}

	return v
}

func randInt(max int) int {
	if max < 1 {
		max = 1
	}

	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if r == nil {
		return 0
	}

	return int(r.Int64())
}

func (c *Client) Words(ctx context.Context, q *WordsQuery) ([]*Word, error) {
	u := &url.URL{
		Path:     "app/words/",
		RawQuery: q.Values().Encode(),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %s", ErrResponseStatus, resp.Status)
	}

	w := map[string]string{}
	err = json.NewDecoder(resp.Body).Decode(&w)
	if err != nil {
		return nil, err
	}

	words := make([]*Word, 0, len(w))
	for word, imageID := range w {
		words = append(words, &Word{
			Word:    word,
			ImageID: imageID,
		})
	}

	return words, nil
}
