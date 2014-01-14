package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"    // Working at 2002271f2160a4d243f0308af0827893e2868157
	"github.com/darkhelmet/twitterstream" // Working at 4051c41877496d38d54647c35897e768fd34385f
	"io/ioutil"
	"log"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("hi")
	b, e := ioutil.ReadFile("./twittercfg")
	if e != nil {
		log.Fatal("Could not read the ./twittercfg file.")
	}
	twittertemp := string(b)
	twitterbits := strings.Split(twittertemp, "\n")
	if len(twitterbits) != 5 {
		log.Fatal("Not enought things in twitter cfg, Needs to be (seperated by \\n) username, consumerKey, consumerSecret, accessToken, accessSecret")
	}
	Client := twitterstream.NewClient(twitterbits[1], twitterbits[2], twitterbits[3], twitterbits[4])
	Conn, e := Client.Track(fmt.Sprintf("@%s", twitterbits[0]))
	// Streamign API is setup now, now just setup the general purpose one now
	anaconda.SetConsumerKey(twitterbits[1])
	anaconda.SetConsumerSecret(twitterbits[2])
	api := anaconda.NewTwitterApi(twitterbits[3], twitterbits[4])

	if e != nil {
		log.Fatal("could not open a streaming connection to get mentions :(")
	}
	for {
		t, e := Conn.Next()
		if e == nil {
			fmt.Printf("TWEET: %s\n", t.Text)
			fmt.Printf("OWNER @%s\n", strings.ToLower(twitterbits[0]))
			if strings.HasPrefix(strings.ToLower(t.Text), fmt.Sprintf("@%s", strings.ToLower(twitterbits[0]))) {
				v := url.Values{} // I dont even know
				t, e := api.PostTweet(fmt.Sprintf("@%s pong", t.User.ScreenName), v)
				if e == nil {
					fmt.Println(t)
				} else {
					fmt.Println(e)
				}
			} else {
				fmt.Println("Does not start with @<user> ignoring")
			}
		}
	}

}
