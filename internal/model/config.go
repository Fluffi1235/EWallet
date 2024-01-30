package model

type Config struct {
	ServicePort string `yaml:"port"`
	DataDB      string `yaml:"dataDB"`
	TimeoutTx   int    `yaml:"timeoutTx"`
}
