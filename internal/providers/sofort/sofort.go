package sofort

import (
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"log"
	"mockgo/internal/config"
)

type SofortProvider struct {
	url string
	notificationUrl string
	notificationTime int
}

func NewProvider(configuration config.Config) SofortProvider {
	return SofortProvider{
		url: configuration.Sofort.Url,
		notificationUrl: configuration.Sofort.NotificationUrl,
		notificationTime: configuration.Sofort.NotificationTime,
	}
}

func (s SofortProvider) GetUrl() string {
	return s.url
}

func (s SofortProvider) Handler(w http.ResponseWriter, r *http.Request) {
	userVariable := r.URL.Query().Get("suvar1")

	if userVariable == "" {
		fmt.Fprint(w, "Sofort Provider received no user variable")
	}

	go func() {
		time.Sleep(time.Duration(s.notificationTime) * time.Second)
		//no error handling for now
		resp, _ := http.Get(s.notificationUrl);

		bodyBytes, _ := ioutil.ReadAll(resp.Body)

		bodyString := string(bodyBytes)
		log.Printf("Received body: %s", bodyString)
	}()

	fmt.Fprintf(w, "Sofort Provider received user variable 1: %s. Sending notification", userVariable)
}
