package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	ChainNormal  ChainNormal  `yaml:"chain_normal"`
	ChainSpecial ChainSpecial `yaml:"chain_special"`
}

type ChainNormal struct {
	ChainTypeA string `yaml:"chain_A"`
	ChainTypeB string `yaml:"chain_B"`
	ChainTypeC string `yaml:"chain_C"`
	ChainTypeD string `yaml:"chain_D"`
	ChainTypeE string `yaml:"chain_E"`
	ChainTypeF string `yaml:"chain_F"`
	ChainTypeG string `yaml:"chain_G"`
	ChainTypeH string `yaml:"chain_H"`
	ChainTypeI string `yaml:"chain_I"`
	ChainTypeJ string `yaml:"chain_J"`
	ChainTypeK string `yaml:"chain_K"`
	ChainTypeL string `yaml:"chain_L"`
	ChainTypeM string `yaml:"chain_M"`
}

type ChainSpecial struct {
	ChainTypeS1 string `yaml:"chain_S1"`
	ChainTypeS2 string `yaml:"chain_S2"`
}

func GetConfigData() string {
	configFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatalf("File reading error: %v", err)
	}
	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err)
	}
	chain := config.ChainNormal.ChainTypeA
	return chain
}

//
//func main() {
//	data := GetConfigData()
//	fmt.Println(data)
//}
