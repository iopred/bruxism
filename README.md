# Bruxism
A chat bot for YouTube, Discord and IRC.

[![GoDoc](https://godoc.org/github.com/iopred/bruxism?status.svg)](https://godoc.org/github.com/iopred/bruxism)
[![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/iopred/bruxism)
[![Go report](http://goreportcard.com/badge/iopred/bruxism)](http://goreportcard.com/report/iopred/bruxism)
[![Build Status](https://travis-ci.org/iopred/bruxism.svg?branch=master)](https://travis-ci.org/iopred/bruxism)

## Current support:

* `!comic [<1-6>]` - Creates a comic from recent messages.
* `!customcomic [<id>:] <text> | [<id>:] <text>` - Creates a custom comic.
* `!customcomicsimple [<id>:] <text> | [<id>:] <text>` - Creates a simple custom comic.
* `!help [<topic>]` - Returns generic help or help for a specific topic. Available topics: `comic,remind`
* `!invite <id>` - Joins the provided YouTube chat, IRC channel or Discord server.
* `!numbertrivia [<number>]` - Returns trivia for a random number or a specified number if provided.
* `!playing` - Set which game the bot is playing on Discord.
* `!reminder <time> | <reminder>` - Sets a reminder.
* `!stats` - Lists bot statistics.
* `!streamer <streamername|streamerid>` - Grabs details about a YouTube streamer.
* `!topstreamers` - List the current top streamers on YouTube Gaming.

## Arguments:

* `youtubeurl` - Outputs a new OAuth URL for YouTube and then exits.
* `youtubeauth` - Sets the YouTube OAuth Response, this will generate a OAuth token.
* `youtubeconfig` - The filename for your YouTube OAuth client. (Download JSON in Google Developers Console -> Credentials).
* `youtubetoken` - The filename for your OAuth token.
* `youtubelivechatids` - An additional comma separated list of YouTube Live Chat ID's to listen on. By default the bot listens to all the live broadcasts on the primary account.
* `discordemail` - Sets the Discord email address.
* `discordpassword` - Sets the Discord password.
* `imgurid` - Sets the Imgur client id

eg:

`bruxism -youtubelivechatids <comma separated list of chat ids> -discordemail <discord email> -discordpassword <discord password> -imgurid <imgur id>`
