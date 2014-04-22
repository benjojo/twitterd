package main

import (
	"bytes"
	"fmt"
	"github.com/ChimeraCoder/anaconda"    // Working at 2002271f2160a4d243f0308af0827893e2868157
	"github.com/darkhelmet/twitterstream" // Working at 4051c41877496d38d54647c35897e768fd34385f
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func IsFile(path string) bool {
	f, e := os.Stat(path)
	if e != nil {
		return false
	}
	if f.IsDir() {
		return false
	}
	return true
}

func CheckForCGIDir() {
	f, e := os.Stat("./cgi")
	if e == nil {
		if !f.IsDir() {
			Logger.Println(`So you have made a cgi file. not a directory.\n
			 What. Removing your sillyness and doing it the right way`)
			e = os.Remove("./cgi")
			if e != nil {
				Logger.Fatal("Cannot remove (silly) the cgi file. What have you done!? (Permission probs)")
			}
			e := os.Mkdir("./cgi", 600)
			if e != nil {
				Logger.Fatalf("Cannot create the cgi dir. I kinda need to stop now. Reason %s", e.Error())
			}
		}
	} else {
		e := os.Mkdir("./cgi", 600)
		if e != nil {
			Logger.Fatalf("Cannot create the cgi dir. I kinda need to stop now. Reason %s", e.Error())
		}
	}
}

func LaunchARealReply(path string, tweet *twitterstream.Tweet, api *anaconda.TwitterApi, ackwithfav bool) {
	cmd := exec.Command(path + getprefix())
	cmd.Env = []string{
		fmt.Sprintf("tweet_text=%s", tweet.Text),
		fmt.Sprintf("tweet_id=%d", tweet.Id),
		fmt.Sprintf("tweet_src=%s", tweet.User.ScreenName),
		fmt.Sprintf("tweet_src_nomention=%s", strings.Join(strings.Split(tweet.Text, " ")[1:], " ")),
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		Logger.Printf("Error launching CGI to serve tweet: Error: %s", err)
	} else {
		if out.String() != "" {
			v := url.Values{} // I dont even know
			v.Add("in_reply_to_status_id", fmt.Sprintf("%d", tweet.Id))
			api.PostTweet(fmt.Sprintf("@%s %s", tweet.User.ScreenName, out.String()), v)
			Logger.Printf("Tweet came in, Replied with %s", fmt.Sprintf("@%s %s", tweet.User.ScreenName, out.String()))
		} else {
			Logger.Println("Empty responce from CGI script. Not sending a blank tweet")
		}
		if ackwithfav {
			api.Favorite(tweet.Id)
		}
	}
}

func LaunchReply(tweet *twitterstream.Tweet, api *anaconda.TwitterApi, ackwithfav bool) {
	if IsFile("./cgi/reply" + getprefix()) {
		LaunchARealReply("./cgi/reply", tweet, api, ackwithfav)
	} else {
		Logger.Println("Reply script does not exist. Try making one?")
	}

	if IsFile("./cgi/reply_" + tweet.User.ScreenName + getprefix()) {
		LaunchARealReply("./cgi/reply_"+tweet.User.ScreenName, tweet, api, ackwithfav)
	}

}

func LaunchMention(tweet *twitterstream.Tweet, api *anaconda.TwitterApi, reply bool) {
	if IsFile("./cgi/mention" + getprefix()) {
		cmd := exec.Command("./cgi/mention" + getprefix())
		cmd.Env = []string{
			fmt.Sprintf("tweet_text=%s", tweet.Text),
			fmt.Sprintf("tweet_id=%d", tweet.Id),
			fmt.Sprintf("tweet_src=%s", tweet.User.ScreenName),
			fmt.Sprintf("tweet_src_name=%s", tweet.User.Name),
			fmt.Sprintf("tweet_src_followers=%s", tweet.User.FollowersCount),
			fmt.Sprintf("tweet_src_nomention=%s", strings.Join(strings.Split(tweet.Text, " ")[1:], " ")),
		}
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			Logger.Printf("Error launching CGI to serve tweet: Error: %s", err)
		} else {
			if reply {
				if out.String() != "" {
					v := url.Values{} // I dont even know
					v.Add("in_reply_to_status_id", fmt.Sprintf("%d", tweet.Id))
					api.PostTweet(fmt.Sprintf("@%s %s", tweet.User.ScreenName, out.String()), v)
					Logger.Printf("Tweet came in, Replied with %s", fmt.Sprintf("@%s %s", tweet.User.ScreenName, out.String()))
				} else {
					Logger.Println("CGI responce was empty. Not sending a blank tweet.")
				}
			}
		}
	} else {
		Logger.Println("Mention script does not exist. Try making one?")
	}

}

func getprefix() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}
	return ""
}
