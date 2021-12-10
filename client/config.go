package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

var ErrServiceNotFound = errors.New("service not found")
var ycp  ConfigProvider
type ConfigProvider interface {
	GetServiceConfig(servicename string) (*Config, error)
}
type Config struct {
	Endpoint string
}

type yamlConfigProvider struct {
	Services map[string]*Config `json:"services"`
}

func (y *yamlConfigProvider) GetServiceConfig(servicename string) (*Config, error) {
	ycp, ok := y.Services[servicename]
	if !ok {
		return nil, ErrServiceNotFound
	}
	return ycp, nil
}
func NewYamlConfigProvider(filepath string) (*yamlConfigProvider, error) {
	content, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	data, _ := ioutil.ReadAll(content)
	ycp := &yamlConfigProvider{}
	_ = json.Unmarshal(data, ycp)

	return ycp,nil
}

func init() {
	ycp,_ = NewYamlConfigProvider("./application.json")
}
