package main

import (
	"net/http/httptest"
	"testing"

	"github.com/beme/abide"
)

func TestRequests(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	// abide
	res := w.Result()
	abide.AssertHTTPResponse(t, "gotham go demo", res)
}
