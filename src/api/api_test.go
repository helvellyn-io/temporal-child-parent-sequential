package api

import (
	"net/http"
	"testing"
)

func TestAdd(t *testing.T) {

	url, _ := http.Get("https://api.openweathermap.org/data/2.5/weather?q=Boulder&appid=050224087da57dabdc13099b40e747e0&units=imperial")
	got := url.StatusCode
	want := 200
	if got != want {
		t.Errorf("got HTTP status %q, wanted HTTP status %q", got, want)
	}
}
