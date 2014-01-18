twitterd
========

Make twitter bots with speed and ease with a CGI like interface.

##Setup

The configuration files are as follows:
```json
{
  "Username": "Fillmein",
  "ConsumerKey": "Fillmein",
  "ConsumerSecret": "Fillmein",
  "AccessToken": "Fillmein",
  "AccessSecret": "Fillmein",
  "StreamTarget": "default",
  "EnableReply": true,
  "EnableMention": false,
  "EnableReplyMention": false
}
```

`Username` is used in checking if the target account was @'td
 
`ConsumerKey` `ConsumerSecret` `AccessToken` `AccessSecret` are twitter API details. You **must** fill these in for the application to work.

`EnableReply` Enables the CGI script `./cgi/reply` to be exectuted when the account is @'td

`EnableMention` Enables the CGI script `./cgi/mention` to be exectuted when anything appeares in the stream

`StreamTarget` Is what the target stream search is, by default the value is "default" meaning that it will look for "@username".

##Getting started

When you run the system for the first time. A config file and CGI folder will be made. 

Depending on what you have enabled in the config settings, by default `EnableReply` is enabled.

This means if your account is named `@testbot` if anyone tweets `@testbot hello world` the `./cgi/reply` script will be exectuted.

This can be a shell script or what ever you want.

the infomation of the tweet can be found in the env vars of the script.

If you where to do "export" while in the CGI script you will see somthing like this

```
declare -x OLDPWD
declare -x PWD="/tmp/twitterd"
declare -x SHLVL="1"
declare -x tweet_id="424612189752659969"
declare -x tweet_src="benjojo"
declare -x tweet_src_nomention="hello world"
declare -x tweet_text="@testbot hello world"
``

For this example we will reply back with the tweetID.

We can do this by putting the following in the script:

```shell
#!/bin/bash
echo -n "Served by twitterd: "
echo -n $tweet_id
```

Note how we didnt need to put @username in there. Twitterd already handles that for you.

And there we go, You now have a simple twitter bot running.

