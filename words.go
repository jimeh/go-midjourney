package midjourney

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
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

func (wq *WordsQuery) URLValues() url.Values {
	v := url.Values{}
	if wq.Query != "" {
		v.Set("query", wq.Query)
	}
	if wq.Amount != 0 {
		v.Set("amount", strconv.Itoa(wq.Amount))
	}
	v.Set("page", strconv.Itoa(wq.Page))
	if wq.RandomSeed {
		v.Set("seed", strconv.Itoa(randInt(9999)))
	} else if wq.Seed != 0 {
		v.Set("seed", strconv.Itoa(wq.Seed))
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
	w := map[string]string{}
	err := c.API.Get(ctx, "app/words/", q.URLValues(), &w)
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
