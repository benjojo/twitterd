package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"    // Working at 2002271f2160a4d243f0308af0827893e2868157
	"github.com/darkhelmet/twitterstream" // Working at 4051c41877496d38d54647c35897e768fd34385f
	"log"
	"net/url"
	"strings"
)

func main() {
	log.Println("Twitterd Started")
	tfg := GetCFG()
	Client := twitterstream.NewClient(tfg.ConsumerKey, tfg.ConsumerSecret, tfg.AccessToken, tfg.AccessSecret)
	Conn, e := Client.Track(fmt.Sprintf("@%s", tfg.Username))
	// Streaming API is setup now, now just setup the general purpose one now
	anaconda.SetConsumerKey(tfg.ConsumerKey)
	anaconda.SetConsumerSecret(tfg.ConsumerSecret)
	api := anaconda.NewTwitterApi(tfg.AccessToken, tfg.AccessSecret)

	if e != nil {
		log.Fatal("could not open a streaming connection to get mentions :(")
	}
	for {
		t, e := Conn.Next()
		if e == nil {
			log.Println("TWEET: %s\n", t.Text)
			log.Println("OWNER @%s\n", strings.ToLower(tfg.Username))
			if strings.HasPrefix(strings.ToLower(t.Text), fmt.Sprintf("@%s", strings.ToLower(tfg.Username))) {
				v := url.Values{} // I dont even know
				t, e := api.PostTweet(fmt.Sprintf("@%s pong", t.User.ScreenName), v)
				if e == nil {
					fmt.Println(t)
				} else {
					fmt.Println(e)
				}
			} else {
				log.Println("Does not start with @<user> ignoring")
			}
		}
	}

}
