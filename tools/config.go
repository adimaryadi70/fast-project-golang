package tools

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type ConfigModel struct {
	TYPE_DATABASE      string `json:"TYPE_DATABASE"`
	TOKEN_JWT_DURATION string `json:"TOKEN_JWT_DURATION"`
	TOKEN_SECRET       string `json:"TOKEN_SECRET"`
	PORT               string `json:"PORT"`
	HOST_DATABASE      string `json:"HOST_DATABASE"`
	USER_DATABASE      string `json:"USER_DATABASE"`
	PASSWORD_DATABASE  string `json:"PASSWORD_DATABASE"`
	PORT_DATABASE      string `json:"PORT_DATABASE"`
	DATABASE           string `json:"DATABASE"`
}

func SetConfig() {
	// configuration support mysql and postgres
	// DB value = mysql || postgres
	configData, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("Failed Read File:", err)
	}
	var config ConfigModel
	err = json.Unmarshal(configData, &config)
	if err != nil {
		log.Fatal("Error Unmarshal:", err)
	}

	os.Setenv("DB", config.TYPE_DATABASE)
	os.Setenv("token_duration", config.TOKEN_JWT_DURATION)
	os.Setenv("token_secret", config.TOKEN_SECRET)
	os.Setenv("portRun", config.PORT)
	os.Setenv("host", config.HOST_DATABASE)
	os.Setenv("user", config.USER_DATABASE)
	os.Setenv("pass", config.PASSWORD_DATABASE)
	os.Setenv("portDB", config.PORT_DATABASE)
	os.Setenv("database", config.DATABASE)
}
