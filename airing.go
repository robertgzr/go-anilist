package anilist

import (
	"time"
)

type Airing struct {
	RawTime     string `json:"time"`
	Countdown   int    `json:"countdown"`
	NextEpisode int    `json:"next_episode"`
}

func (airing *Airing) Time() (t time.Time, err error) {
	t, err = time.Parse(time.RFC3339, airing.RawTime)
	return
}
