package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRouter(t *testing.T) {
	start := time.Now()
	r := Router()
	elapsed := time.Since(start).Seconds()

	if elapsed <= 35 {
		t.Errorf("Start up duration too short, got: %f", elapsed)
	}

	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code for home path, got: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	res, err = http.Get(ts.URL + "/healthz")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code for healthz endpoint, got: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	res, err = http.Get(ts.URL + "/greetings")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Unexpected status code for greetings endpoint, got: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	res, err = http.Get(ts.URL + "/nonexistent")
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Unexpected status code for /nonexistent, got: %d, want: %d.", res.StatusCode, http.StatusNotFound)
	}
}
