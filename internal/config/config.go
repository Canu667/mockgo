package config

type Config struct {
	Dummy `yaml:"dummy"`
	Sofort `yaml:"sofort"`
}

type Dummy struct {
	Url string `yaml:"url"`
}

type Sofort struct {
	Url              string `yaml:"url"`
	NotificationUrl  string `yaml:"notificationUrl"`
	NotificationTime int `yaml:"notificationTime"`
}