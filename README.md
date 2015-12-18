# Bruxism
A chat bot for YouTube and Discord.

Current support:

* Slow Mode - Slows down YouTube chat by temporarily banning anyone who speaks.
* Top Streamers - Lists the current top streamers on YouTube Gaming.
* Help - !help to list all available commands.

Arguments:

* youtubeurl - Outputs a new OAuth URL for YouTube and then exits.
* youtubeauth - Sets the YouTube OAuth Response, this will generate a OAuth token.
* youtubeconfig - The filename for your YouTube OAuth client. (Download JSON in Google Developers Console -> Credentials).
* youtubetoken - The filename for your OAuth token.
* youtubelivechatids - An additional comma separated list of YouTube Live Chat ID's to listen on. By default the bot listens to all the live broadcasts on the primary account.
* discordemail - Sets the Discord email address.
* discordpassword - Sets the Discord password.

eg:

```
bruxism -youtubelivechatids <comma separated list of chat ids> -discordemail <discord email> -discordpassword <discord password>
```
