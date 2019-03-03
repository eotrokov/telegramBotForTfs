package config

import (
	"github.com/tkanos/gonfig"
	"log"
)

type Configuration struct {
	BotToken string
	TfsLogin string
	TfsToken string
	TfsUrl string
}

func GetConfig() Configuration {
	configuration := Configuration{}
	err := gonfig.GetConf("config/config.json", &configuration)
	if err != nil {
		log.Panic(err)
	}
	return configuration
}