package main

import (
	"flag"
	"log"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"net/http"
	"gopkg.in/yaml.v2"
	"mockgo/internal/providers"
	"mockgo/internal/config"
)

func main() {
	filename, _ := filepath.Abs("./internal/config/providers.yaml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	configYaml := config.Config{}

	err = yaml.Unmarshal(yamlFile, &configYaml)
	if err != nil {
		panic(err)
	}

	port := flag.Int("p", 8080, "Port on which the service will work")
	flag.Parse()

	for _, provider := range providers.GetProviders(configYaml) {
		log.Printf(provider.GetUrl());
		http.HandleFunc(provider.GetUrl(), provider.Handler)
	}

	log.Printf("Server starting on port %v\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", *port), nil))
}