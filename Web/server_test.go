package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWebServer(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	WebServer(response, request)

	t.Run("Retourne le site web", func(t *testing.T) {
		got := response.Body.String()
		want := "Hello"

		if got != want {
			t.Errorf("On a %q, On voudrait %q", got, want)
		}
	})

}