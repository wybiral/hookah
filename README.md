# hookah
hookah is a lightweight pipeline tool for orchestrating WebSocket streams.

Operations like fanning-in multiple data sources, fanning-out to multiple
consumers, and applying local map/filter operations between endpoints are all
possible using hookah and standard POSIX pipes.

### Install dependencies
```
go get github.com/gorilla/websocket
go get github.com/wybiral/hookah/cmd/hookah
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
