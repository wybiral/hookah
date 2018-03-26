# hookah
hookah is a lightweight pipeline tool that makes it easy to work with WebSocket
streams as though they were normal POSIX-style pipelines.

The goal is to makes it easy to perform tasks like fanning-in multiple data
sources to create a firehose API or fanning-out a local program's output for
real-time distribution. All without having to integrate any WebSocket-specific
code into your applications.

### Install dependencies
```
go get github.com/gorilla/websocket
go get github.com/wybiral/hookah
```
### Build
```
go build github.com/wybiral/hookah/cmd/hookah
```
### Usage
To see usage instructions run:
```
hookah
```
To start a relay node:
```
hookah node host:port
```
Pipe a local program into a remote node:
```
local_program | hookah send remote_host:port
```
Read from a remote node:
```
hookah recv remote_host:port
```
Pipe between two remote nodes and apply local program as map/filter:
```
hookah recv remote_host1:port | local_program | hookah send remote_host2:port
```
