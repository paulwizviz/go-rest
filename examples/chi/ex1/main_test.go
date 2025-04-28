package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	r := router()
	r.ServeHTTP(rec, req)
	got := rec.Body.Bytes()
	want := []byte("welcome")
	if !bytes.Equal(got, want) {
		t.Fatalf("Want: %s Got: %s", want, got)
	}
}
