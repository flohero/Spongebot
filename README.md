# Spongebot
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)
[![Build Status](https://travis-ci.org/flohero/Spongebot.svg?branch=master)](https://travis-ci.org/flohero/Spongebot)

Add custom commands to this bot, via a simple REST API.

**Not finished and not tested**

## Scripting
You can create custom functionality through 
[starlark](https://docs.bazel.build/versions/master/skylark/language.html), 
a python dialect and [starlight-go](https://github.com/starlight-go/starlight). 
It's simple to use, the program provides a struct called `s` which contains two fields, 
a `Message` field which contains the discord message and a `Result` field, 
in which you should save the final answer. 

### Example
```python
s.Result = s.Message.upper()
```
This script will return the message in uppercase.
