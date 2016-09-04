# agentstats
A simple wrapper for the api of https://www.agent-stats.com as documented at http://apidocs.agent-stats.com/

# Getting started
First of all, get your api key which you can generate or find at https://www.agent-stats.com/preferences.php
After that you can get started with this library.

To list all the groups you're in you can do something like the following.
```go
client := agentstats.NewClient("<api token>")

groups, err := client.Groups()
```
Where groups will have all the groups that you're in, each group has an ID which you can later use to call GroupProgress(string) like the following.
```go
client := agentstats.NewClient("<api token>")

progress, err := client.GroupProgress("<group ID>")
```
