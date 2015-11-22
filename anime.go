package anilist

import (
	"encoding/json"
)

type AnimeSmall struct {
	Id   int    `json:"id"`
	Type string `json:"type"`

	AiringStatus  string `json:"airing_status"`
	TotalEpisodes int    `json:"total_episodes"`

	TitleRomaji   string   `json:"title_romaji"`
	TitleJapanese string   `json:"title_japanese"`
	TitleEnglish  string   `json:"title_english"`
	Synonyms      []string `json:"synonyms"`

	ImageURLSmall  string `json:"image_url_sml"`
	ImageURLMedium string `json:"image_url_med"`
	ImageURLLarge  string `json:"image_url_lge"`

	RelationType string `json:"relation_type"`
}

type Anime struct {
	Id   int    `json:"id"`
	Type string `json:"type"`

	TitleRomaji   string   `json:"title_romaji"`
	TitleJapanese string   `json:"title_japanese"`
	TitleEnglish  string   `json:"title_english"`
	Synonyms      []string `json:"synonyms"`
	Genres        []string `json:"genres"`

	Description string `json:"description"`

	ImageURLSmall  string `json:"image_url_sml"`
	ImageURLMedium string `json:"image_url_med"`
	ImageURLLarge  string `json:"image_url_lge"`
	ImageURLBanner string `json:"image_url_banner"`

	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`

	Hashtag string `json:"hashtag"`
	Source  string `json:"source"`

	DurationMinutes int     `json:"duration"`
	AiringStatus    string  `json:"airing_status"`
	TotalEpisodes   int     `json:"total_episodes"`
	Airing          *Airing `json:"airing"`

	RelationType string `json:"relation_type"`
}

func (anime *Anime) AnimeSmall() (sml *AnimeSmall, err error) {
	var (
		rawData []byte
	)

	rawData, err = json.Marshal(anime)
	if err != nil {
		return
	}

	sml = new(AnimeSmall)
	err = json.Unmarshal(rawData, &sml)
	if err != nil {
		return
	}

	return
}
