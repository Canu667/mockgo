package dummy

import (
	"fmt"
	"net/http"
	"mockgo/internal/config"
)

type DummyProvider struct {
	url string
}

func NewProvider(configuration config.Config) DummyProvider {
	return DummyProvider{url: configuration.Dummy.Url}
}

func (d DummyProvider) GetUrl() string {
	return d.url
}

func (d DummyProvider) Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Dummy Provider is working\n")
}
