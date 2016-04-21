=====
This is a music plugin for the bruxism multi-service bot.

This code is in very early stages and is still a bit messy.


### Requirements

* [ffmpeg](http://ffmpeg.org/) must be installed and within your PATH
* [youtube-dl](https://github.com/rg3/youtube-dl) must be installed and exist in the same folder as bruxism.
* [dca](https://github.com/bwmarrin/dca) must be installed and exist in the same folder as bruxism.

### Music Plugin Commands

* NOTE: All commands must be prefixed with a @mention of your bot
* NOTE: `music` can be abbreviated with just `mu`

#### Format

`@botname music <required> [optional] [this | or this]`


#### Commands Quick Reference


| Command              | Arguments                        | Action
| -------------------- | -------------------------------- | -------------------
| **mu**sic **join**   | `[channelID]`                    | Join channel of caller or provided channel ID
| **mu**sic **leave**  |                                  | Leave the currently connected voice channel.
| **mu**sic **play**   | `[url]`                          | Start queue player and optionally enqueue a URL
| **mu**sic **info**   |                                  | Display plugin / queue / playing song info
| **mu**sic **pause**  |                                  | Pause playback.
| **mu**sic **resume** |                                  | Resume playback.
| **mu**sic **skip**   |                                  | Skip current song.
| **mu**sic **stop**   |                                  | Stop queue player.
| **mu**sic **loop**   |                                  | Toggle Loop Queue setting
| **mu**sic **list**   |                                  | List all items in queue.
| **mu**sic **clear**  |                                  | Clear all items from queue.

### Notes

For faster startup of youtube-dl use github cloned version of youtube-dl, see
https://www.raspberrypi.org/forums/viewtopic.php?f=38&t=83763

To avoid youtube-dl ffmpeg post processing error, use a youtube-dl version past
this commit https://github.com/rg3/youtube-dl/commit/e38cafe986994d65230e6518752def8b53ad7308
see issue https://github.com/rg3/youtube-dl/issues/8729#event-574790956

### Author

[Bruce Marriner](https://github.com/bwmarrin/discordgo)
