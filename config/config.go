package config

import (
	"encoding/json"

	"github.com/kyokomi/appConfig"
)

const configFileName = "config.json"

type AccessConfig struct {
	//	Host    string `json:"host"`
	//	ApiPath string `json:"api_path"`
	Token string `json:"token"`
}

var defaultConfig = AccessConfig{
	//	Host:    "https://gitlab.com/",
	//	ApiPath: "api/v3/",
	Token: "aaaaaaaaaaaaaaaaaaaaaaa",
}

type CliAppConfig struct {
	appConfig.AppConfig
	AccessConfig AccessConfig
}

func NewCliAppConfig(appName string) *CliAppConfig {
	return &CliAppConfig{
		AppConfig:    *appConfig.NewAppConfig(appName, configFileName),
		AccessConfig: defaultConfig,
	}
}

func (a *CliAppConfig) ReadAccessTokenJson() error {
	data, err := a.ReadAppConfig()
	if err != nil {
		return err
	}
	var c AccessConfig
	if err := json.Unmarshal(data, &c); err != nil {
		return err
	}
	a.AccessConfig = c

	return nil
}

func (a CliAppConfig) WriteDefaultAccessConfig() error {
	return a.WriteAccessConfig(&defaultConfig)
}

func (a CliAppConfig) WriteAccessConfig(config *AccessConfig) error {
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	return a.WriteAppConfig(data)
}
