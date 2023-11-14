package main

import (
	"fmt"
	"testing"
)

func TestNewGHIssue(t *testing.T) {
	req := GHIssue{
		Title:  "foo",
		Body:   "baz",
		Owner:  "PMaynard",
		Repo:   "NRCBot",
		Labels: []string{"Ideas"},
	}

	res, err := NewGHIssue(req)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
