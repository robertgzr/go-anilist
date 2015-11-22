package anilist

type CharacterSmall struct {
	Id             int    `json:"id"`
	FirstName      string `json:"name_first"`
	LastName       string `json:"name_last"`
	ImageURLMedium string `json:"image_url_med"`
	ImageURLLarge  string `json:"image_url_lge"`
	Role           string `json:"role"`
}

type Character struct {
	Id             int    `json:"id"`
	FirstName      string `json:"name_first"`
	LastName       string `json:"name_last"`
	NameJapanese   string `json:"name_japanese"`
	ImageURLMedium string `json:"image_url_med"`
	ImageURLLarge  string `json:"image_url_lge"`
	Role           string `json:"role"`
}
