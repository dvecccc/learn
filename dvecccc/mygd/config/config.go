package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	ChainNormal  ChainNormal  `yaml:"chainNormal"`
	ChainSpecial ChainSpecial `yaml:"chainSpecial"`
}

type ChainNormal struct {
	ChainTypeA string `yaml:"chainTypeA" json:"chainTypeA"`
	ChainTypeB string `yaml:"chainTypeB" json:"chainTypeB"`
	ChainTypeC string `yaml:"chainTypeC" json:"chainTypeC"`
	ChainTypeD string `yaml:"chainTypeD" json:"chainTypeD"`
	ChainTypeE string `yaml:"chainTypeE" json:"chainTypeE"`
	ChainTypeF string `yaml:"chainTypeF" json:"chainTypeF"`
	ChainTypeG string `yaml:"chainTypeG" json:"chainTypeG"`
	ChainTypeH string `yaml:"chainTypeH" json:"chainTypeH"`
	ChainTypeI string `yaml:"chainTypeI" json:"chainTypeI"`
	ChainTypeJ string `yaml:"chainTypeJ" json:"chainTypeJ"`
	ChainTypeK string `yaml:"chainTypeK" json:"chainTypeK"`
	ChainTypeL string `yaml:"chainTypeL" json:"chainTypeL"`
	ChainTypeM string `yaml:"chainTypeM" json:"chainTypeM"`
}

type ChainSpecial struct {
	ChainTypeSpecialOne string `yaml:"chainSpecialOne" json:"chainTypeSpecialOne"`
	ChainTypeSpecialTwo string `yaml:"chainSpecialTwo" json:"chainTypeSpecialTwo"`
}

func GetConfigData() string {
	configFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatalf("read config file error: %v", err)
	}
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("unmarshal config file error: %v	", err)
	}
	return config.ChainNormal.ChainTypeA
}