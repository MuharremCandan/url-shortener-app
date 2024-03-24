package config

import "github.com/spf13/viper"

type Config struct {
	HttpServer struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"httpserver"`
	Database struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Name string `yaml:"name"`
	} `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
