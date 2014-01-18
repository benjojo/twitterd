package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"    // Working at 2002271f2160a4d243f0308af0827893e2868157
	"github.com/darkhelmet/twitterstream" // Working at 4051c41877496d38d54647c35897e768fd34385f
	"log"
	"strings"
)

func main() {
	log.Println("Twitterd Started")
	tfg := GetCFG()
	Client := twitterstream.NewClient(tfg.ConsumerKey, tfg.ConsumerSecret, tfg.AccessToken, tfg.AccessSecret)
	var Conn *twitterstream.Connection
	var e error

	if tfg.StreamTarget == "default" || tfg.StreamTarget == "" {
		Conn, e = Client.Track(fmt.Sprintf("@%s", tfg.Username))
	} else {
		Conn, e = Client.Track(tfg.StreamTarget)
	}
	// Streaming API is setup now, now just setup the general purpose one now
	anaconda.SetConsumerKey(tfg.ConsumerKey)
	anaconda.SetConsumerSecret(tfg.ConsumerSecret)
	api := anaconda.NewTwitterApi(tfg.AccessToken, tfg.AccessSecret)
	CheckForCGIDir()
	if e != nil {
		log.Fatalf("could not open a streaming connection to get mentions :( Reason: %s \n", e)
	}
	for {
		t, e := Conn.Next()
		if e == nil {
			if strings.HasPrefix(strings.ToLower(t.Text), fmt.Sprintf("@")) { // check if it starts with a @
				// Launch a CGI instance to reply.
				if tfg.EnableReply {
					go LaunchReply(t, api)
				}
			} else {
				if tfg.EnableMention {
					go LaunchMention(t, api, tfg.EnableReplyMention)
				} else {
					log.Println("Does not start with @<user> and since 'EnableMention' is disabled ignoring")
				}
			}
		} else {
			log.Println("I could not poll to get the next tweet. Attempting to reconnect to twitter stream")
			if tfg.StreamTarget == "default" || tfg.StreamTarget == "" {
				Conn, e = Client.Track(fmt.Sprintf("@%s", tfg.Username))
			} else {
				Conn, e = Client.Track(tfg.StreamTarget)
			}
			if e != nil {
				log.Fatal("Could not reconnect to twitter streaming. Exiting")
			}
		}
	}

}
