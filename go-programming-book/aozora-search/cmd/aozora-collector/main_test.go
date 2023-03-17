package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindEntries(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//
	}))
	defer ts.Close()
}
