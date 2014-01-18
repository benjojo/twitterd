package main

import (
	"bytes"
	"fmt"
	"github.com/ChimeraCoder/anaconda"    // Working at 2002271f2160a4d243f0308af0827893e2868157
	"github.com/darkhelmet/twitterstream" // Working at 4051c41877496d38d54647c35897e768fd34385f
	"log"
	"net/url"
	"os"
	"os/exec"
)

func CheckForCGIDir() {
	f, e := os.Stat("./cgi")
	if e == nil {
		if !f.IsDir() {
			log.Println(`So you have made a cgi file. not a directory.\n
			 What. Removing your sillyness and doing it the right way`)
			e = os.Remove("./cgi")
			if e != nil {
				log.Fatal("Cannot remove (silly) the cgi file. What have you done!? (Permission probs)")
			}
			e := os.Mkdir("./cgi", 600)
			if e != nil {
				log.Fatalf("Cannot create the cgi dir. I kinda need to stop now. Reason %s", e.Error())
			}
		}
	} else {
		e := os.Mkdir("./cgi", 600)
		if e != nil {
			log.Fatalf("Cannot create the cgi dir. I kinda need to stop now. Reason %s", e.Error())
		}
	}
}

func LaunchReply(tweet *twitterstream.Tweet, api *anaconda.TwitterApi) {
	cmd := exec.Command("./cgi/reply")
	cmd.Env = append(cmd.Env, fmt.Sprintf("tweet_text:%s", tweet.Text))
	cmd.Env = append(cmd.Env, fmt.Sprintf("tweet_id:%d", tweet.Id))
	cmd.Env = append(cmd.Env, fmt.Sprintf("tweet_src:%s", tweet.User.ScreenName))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Printf("Error launching CGI to serve tweet: Error: %s", err)
	} else {
		v := url.Values{} // I dont even know

		api.PostTweet(fmt.Sprintf("@%s %s", tweet.User.ScreenName, out.String()), v)
		fmt.Printf("in all caps: %q\n", out.String())
	}
}
