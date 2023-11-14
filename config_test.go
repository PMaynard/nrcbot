package main

import (
	"testing"
)

func TestConfig(t *testing.T) {

	config, err := ConfigFile("testdata/config.json")
	if err != nil {
		t.Error(err)
	}

	if config.GithubToken != "github_pat_19283712398712yhej_skdjfheiurh2iruehfd1i192387" {
		t.Error("Config Not Match got: ", config.GithubToken)
	}

}
