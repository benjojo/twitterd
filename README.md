twitterd
========

Make twitter bots with speed and ease with a CGI like interface.

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

`StreamTarget` Is what the target stream search is, by default the value is "default" meaning that it will look for "@<username>".

