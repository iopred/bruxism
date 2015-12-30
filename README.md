# Bruxism
A chat bot for YouTube and Discord.

[![GoDoc](https://godoc.org/github.com/iopred/bruxism?status.svg)](https://godoc.org/github.com/iopred/bruxism)
[![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/iopred/bruxism)
[![Go report](http://goreportcard.com/badge/iopred/bruxism)](http://goreportcard.com/report/iopred/bruxism)
[![Build Status](https://travis-ci.org/iopred/bruxism.svg?branch=master)](https://travis-ci.org/iopred/bruxism)

## Current support:

* Help - !help to list all available commands.
* Invite - !invite to invite the bot into a channel or Discord server.
* Playing - !playing to set the currently playing game in Discord.
* Slow Mode - !slowmode to slows down YouTube chat by temporarily banning anyone who speaks.
* Streamer - !streamer to list stats about a YouTube streamer.
* Top Streamers - !topstreamers to list the current top streamers on YouTube Gaming.

## Arguments:

* youtubeurl - Outputs a new OAuth URL for YouTube and then exits.
* youtubeauth - Sets the YouTube OAuth Response, this will generate a OAuth token.
* youtubeconfig - The filename for your YouTube OAuth client. (Download JSON in Google Developers Console -> Credentials).
* youtubetoken - The filename for your OAuth token.
* youtubelivechatids - An additional comma separated list of YouTube Live Chat ID's to listen on. By default the bot listens to all the live broadcasts on the primary account.
* discordemail - Sets the Discord email address.
* discordpassword - Sets the Discord password.
* imgurid - Sets the Imgur client id

eg:

```
bruxism -youtubelivechatids <comma separated list of chat ids> -discordemail <discord email> -discordpassword <discord password> -imgurid <imgur id>
```
