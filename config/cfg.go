package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Push struct {
		Type   int64 `yaml:"Type"`
		Weixin struct {
			Hook string `yaml:"hook"`
		} `yaml:"Weixin"`
		Telegram struct {
			Token string `yaml:"token"`
		} `yaml:"Telegram"`
		Bark struct {
			Hook string `yaml:"hook"`
		} `yaml:"Bark"`
	} `yaml:"Push"`
	Keys struct {
		BscKey string `yaml:"bsc_key"`
		EthKey string `yaml:"eth_key"`
		SolKey string `yaml:"sol_key"`
	} `yaml:"Keys"`
	Wallet struct {
		Bsc []string `yaml:"bsc"`
		Sol []string `yaml:"sol"`
		Eth []string `yaml:"eth"`
	} `yaml:"Wallet"`
	Exchange struct {
		Binance struct {
			Apikey string `yaml:"apikey"`
			Secret string `yaml:"secret"`
		} `yaml:"binance"`
	} `yaml:"Exchange"`
	Common struct {
		Interval int64 `yaml:"interval"`
		Duration int64 `yaml:"duration"`
	} `yaml:"Common"`
}

func HandleYaml() *Config {
	dir, _ := os.Getwd()
	file, err := os.Open(dir + "/config/config.yaml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("Error decoding YAML:", err)
		return nil
	}
	return &config
}

func ChekcAvaiable() {

}
