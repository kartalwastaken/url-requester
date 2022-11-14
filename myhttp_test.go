package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestNormalizeURL(t *testing.T) {
	url := normalizeURL("google.com")
	if url != "http://google.com" {
		t.Errorf("want: http://google.com, got: %s", url)
	}
	url = normalizeURL("www.google.com")
	if url != "http://www.google.com" {
		t.Errorf("want: www.google.com, got: %s", url)
	}
	url = normalizeURL("test")
	if url != "http://test" {
		t.Errorf("want: http://test, got: %s", url)
	}
}
