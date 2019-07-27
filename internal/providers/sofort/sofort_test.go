package sofort

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"mockgo/internal/config"
)

func TestHandler(t *testing.T) {
	config := config.Config{
		Dummy: config.Dummy {
			Url:"/",
		},
		Sofort: config.Sofort {
			Url: "/",
			NotificationTime: 0,
			NotificationUrl: "http://example.com",
		},
	}
	sofort := NewProvider(config)

	//bonus test for url
	got := sofort.GetUrl()
	want := "/"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

	req, err := http.NewRequest("GET", "/?suvar1=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	sofort.Handler(rr, req);

	expected := "Sofort Provider received user variable 1: 1. Sending notification"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got '%v' want '%v'",
			rr.Body.String(), expected)
	}
}