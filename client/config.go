package main

import (
	"errors"
)

var ErrServiceNotFound = errors.New("service not found")

type ConfigProvider interface {
	GetServiceConfig(servicename string) (*Config, error)
}
type Config struct {
	Endpoint string
}

var icp ConfigProvider

type InMemoryConfigProvider struct {
	Services map[string]*Config
}

func (i *InMemoryConfigProvider) GetServiceConfig(servicename string) (*Config, error) {
	cfg, ok := i.Services[servicename]
	if !ok {
		return nil, ErrServiceNotFound
	}
	return cfg, nil
}
func NewInMemoryConfigProvider() *InMemoryConfigProvider {
	return &InMemoryConfigProvider{
		Services: map[string]*Config{
			"hello": &Config{
				Endpoint: "http://localhost:9000/",
			},
			"redis": &Config{
				Endpoint: "http://localhost:6379",
			},
		},
	}
}

/*
var ycp  ConfigProvider
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
*/
func init() {
	//ycp,_ = NewYamlConfigProvider("./application.json")

	icp = NewInMemoryConfigProvider()
}
