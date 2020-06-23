package config

import "github.com/spf13/viper"

func NewConfig(filePath, typeName string) (*viper.Viper, error) {
	cfg := viper.New()
	cfg.SetConfigName("config")
	cfg.SetConfigType(typeName)
	cfg.AddConfigPath(filePath)
	if err := cfg.ReadInConfig(); err != nil {
		return nil, err
	}
	return cfg, nil
}
