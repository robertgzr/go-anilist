package anilist_test

import (
	"os"
	"testing"

	"github.com/KuroiKitsu/go-anilist"
)

type credentials struct {
	clientId     string
	clientSecret string
}

func TestSession_SearchAnime(t *testing.T) {
	c := getCredentials(t)

	s, err := anilist.NewSession(c.clientId, c.clientSecret)
	if err != nil {
		t.Fatal(err)
	}

	if err := s.Auth(); err != nil {
		t.Fatal(err)
	}

	results, err := s.SearchAnime("shakugan no shana")
	if err != nil {
		t.Fatal(err)
	}

	if len(results) == 0 {
		t.Fatal("no results")
	}
}

func TestSession_SearchCharacter(t *testing.T) {
	s := getSession(t)

	results, err := s.SearchCharacter("kurumi tokisaki")
	if err != nil {
		t.Fatal(err)
	}

	if len(results) == 0 {
		t.Fatal("no results")
	}
}

func getSession(t *testing.T) *anilist.Session {
	c := getCredentials(t)

	s, err := anilist.NewSession(c.clientId, c.clientSecret)
	if err != nil {
		t.Fatal(err)
	}

	if err := s.Auth(); err != nil {
		t.Fatal(err)
	}

	return s
}

func getCredentials(t *testing.T) *credentials {
	clientId := os.Getenv("ANILIST_CLIENT_ID")
	clientSecret := os.Getenv("ANILIST_CLIENT_SECRET")

	if clientId == "" {
		t.Skip("ANILIST_CLIENT_ID environment variable not set")
	}

	if clientSecret == "" {
		t.Skip("ANILIST_CLIENT_SECRET environment variable not set")
	}

	return &credentials{
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}
