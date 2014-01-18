package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type TwitConfig struct {
	Username           string
	ConsumerKey        string
	ConsumerSecret     string
	AccessToken        string
	AccessSecret       string
	StreamTarget       string // In the case you want to choose somthing else to stream that isnt @<username>
	EnableReply        bool   // Do you want to enable /cgi/reply
	EnableMention      bool   // Do you want to enable /cgi/mention
	EnableReplyMention bool   // Do you want the /cgi/mention to reply to things?
}

func GetCFG() TwitConfig {
	b, e := ioutil.ReadFile("./.twittercfg.json")
	var tfg TwitConfig
	if e != nil {
		tfg.AccessSecret = "Fillmein"
		tfg.AccessToken = "Fillmein"
		tfg.ConsumerKey = "Fillmein"
		tfg.ConsumerSecret = "Fillmein"
		tfg.Username = "Fillmein"
		tfg.StreamTarget = "default"
		tfg.EnableReply = true
		out, e := json.Marshal(tfg)
		e = ioutil.WriteFile("./.twittercfg.json", out, 600)
		if e != nil {
			log.Fatal("cannot open settings file :(")
		}
		log.Fatal("Built config file. please fill it in.")
	}

	json.Unmarshal(b, &tfg)
	if tfg.AccessSecret == "Fillmein" || tfg.AccessToken == "Fillmein" || tfg.ConsumerKey == "Fillmein" || tfg.ConsumerSecret == "Fillmein" || tfg.Username == "Fillmein" {
		log.Fatal("You need to fill in the config settings in ./.twittercfg.json")
	}
	return tfg
}
