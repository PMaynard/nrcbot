package main

import (
	"testing"
)

func TestGetDOI(t *testing.T) {
	res, err := GetDOI("10.5555/12345678")
	if err != nil {
		t.Error(err)
	}
	if res.Message.Title[0] != "Toward a Unified Theory of High-Energy Metaphysics: Silly String Theory" {
		t.Error("Title Does not match. Got", res.Message.Title)
	}
}
