package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type TwitConfig struct {
	Username       string
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

func GetCFG() TwitConfig {
	b, e := ioutil.ReadFile("./.twittercfg.json")
	if e != nil {
		log.Fatal("cannot open settings file :(")
	}
	tfg := TwitConfig{}
	json.Unmarshal(b, &tfg)
	return TwitConfig
}
