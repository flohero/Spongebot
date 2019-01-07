# Spongebot
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)
[![Build Status](https://travis-ci.org/flohero/Spongebot.svg?branch=master)](https://travis-ci.org/flohero/Spongebot)

Manage a discord bot via a REST API and an Angular Website.

**Not finished and not tested**

- [Spongebot](#spongebot)
- [Why](#why)
- [Build](#build)
- [Managing the bot](#managing-the-bot)
- [Commands Usage](#commands-usage)
  - [Scripting](#scripting)
    - [The `s` struct and builtin function](#the-s-struct-and-builtin-function)
    - [Example](#example)
- [ToDo](#todo)

# Why
I hated to manage my bot just over source files or via command line, so i made a simple REST API and website to manage my bot.

# Build
To run this Project you need a postgresql instance and a webserver like apache2.

1. Build the go app, with go modules enabled
2. Then rename `.env.sample` to `.env` and add the required informations, like postgresql host, port, discord bot token, etc.
3. The angular website is located in the `website` directory
4. Before you can build the website you have to add the host entry in `website/environment/environment.prod.ts`
5. Then build the website with `npm run-script "build prod"`
6. After that copy the content of `website/dist/website` to your webserver
7. Now you can run the go app you build at step 1
8. The bot should come online on your discord server

> NOTE: Right now the bot needs admin rights on the discord server to properply function.

# Managing the bot
If the bot and website are running, you can manage the bot. The standard login credentials are `sponge` : `bot`. **You should change them immediatly.**

# Commands Usage
Commands are regular expression, which means every message will be matched against all commands. This means a message can match more then one time, in this case the bot will execute all of them.

The can reply with simple static responses or you can build your own reply via scripts.

## Scripting
You can create custom functionality through 
[starlark](https://docs.bazel.build/versions/master/skylark/language.html), 
a python dialect and [starlight-go](https://github.com/starlight-go/starlight). 
It's simple to use, the program provides a struct called `s` which contains two fields, 

### The `s` struct and builtin functions
With s you can call builtin variables. You call them, like in many programming languages, with `s.`.

| Name         |                                                                                                                             Type |
| ------------ | -------------------------------------------------------------------------------------------------------------------------------: |
| `s.Message`  |                                                                                             A string, which contains the message |
| `s.Result`   |                                                                                A empty string, which sould contain your response |
| `s.GuildId`  |                                                                             A string which contains the ID of the discord server |
| `s.AuthorId` |                                                                             A string which contains the ID of the message author |
| `kickUser`   | A function which takes three arguments: a GuildID, AuthorID and a Reasons, all three are strings. Its purpose is to kick an user |

### Example
```python
s.Result = s.Message.upper()
```
This script will return the message in uppercase.

# ToDo
- [ ] Dockerize the app
- [ ] Better tests
- [ ] Make more discord commands available, like ban, votekick, etc
- [x] Better README.md
- [x] Host the angular website with go http
- [ ] Check required rights of the bot on the discord server