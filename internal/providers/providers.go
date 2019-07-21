package providers

import (
	"net/http"
	"mockgo/internal/providers/dummy"
	"mockgo/internal/config"
)

type Provider interface {
	GetUrl() string
	Handler(w http.ResponseWriter, r *http.Request)
}

func GetProviders(config config.Config) []Provider {
	return []Provider{
		dummy.NewProvider(config),
	}
}