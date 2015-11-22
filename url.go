package anilist

import (
	"net/url"
)

var (
	HomeURL *url.URL
	APIURL  *url.URL
)

func initURL() {
	var (
		err error
	)

	HomeURL, err = url.Parse("https://anilist.co")
	if err != nil {
		panic(err)
	}

	APIURL, err = HomeURL.Parse("/api/")
	if err != nil {
		panic(err)
	}
}
