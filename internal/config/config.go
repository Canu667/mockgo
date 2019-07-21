package config

type Config struct {
	Dummy struct {
		Url string `yaml:"url"`
	} `yaml:"dummy"`
}