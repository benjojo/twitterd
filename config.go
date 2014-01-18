package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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

func GetDefaultConfig() TwitConfig {
	var tfg TwitConfig
	tfg.AccessSecret = "Fillmein"
	tfg.AccessToken = "Fillmein"
	tfg.ConsumerKey = "Fillmein"
	tfg.ConsumerSecret = "Fillmein"
	tfg.Username = "Fillmein"
	tfg.StreamTarget = "default"
	tfg.EnableReply = true
	tfg.EnableMention = false
	tfg.EnableReplyMention = false
	return tfg
}

func CheckIfResetConfig(args []string) {
	if len(args) == 2 {
		if args[1] == "reset" {
			e := os.Remove("./.twittercfg.json")
			if e != nil {
				log.Fatal("Could not remove current config file. Permissions issue?")
			}
			Default := GetDefaultConfig()
			out, e := json.Marshal(Default)
			e = ioutil.WriteFile("./.twittercfg.json", out, 600)
			if e != nil {
				log.Fatal("cannot open settings file :(")
			}
			log.Fatal("Built config file. please fill it in.")
		}
	}
}

func GetCFG() TwitConfig {
	b, e := ioutil.ReadFile("./.twittercfg.json")
	tfg := GetDefaultConfig()
	if e != nil {
		out, e := json.Marshal(tfg)
		e = ioutil.WriteFile("./.twittercfg.json", out, 600)
		if e != nil {
			log.Fatal("cannot open settings file :(")
		}
		log.Fatal("Built config file. please fill it in.")
	}

	e = json.Unmarshal(b, &tfg)
	if e != nil {
		log.Fatalf("Could not parse config settings. You can reset the cfg by doing $twitterd reset")
	}
	if tfg.AccessSecret == "Fillmein" || tfg.AccessToken == "Fillmein" || tfg.ConsumerKey == "Fillmein" || tfg.ConsumerSecret == "Fillmein" || tfg.Username == "Fillmein" {
		log.Fatal("You need to fill in the config settings in ./.twittercfg.json")
	}
	return tfg
}
