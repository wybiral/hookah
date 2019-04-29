# hookah
A Swiss Army knife for data pipelines.

[![GoDoc](https://godoc.org/github.com/wybiral/hookah?status.svg)](https://godoc.org/github.com/wybiral/hookah)

# Download options

- [Releases](https://github.com/wybiral/hookah/releases)
- [Sourcecode](https://github.com/wybiral/hookah/archive/master.zip)

# Build instructions
### go get
```
go get github.com/wybiral/hookah/cmd/hookah
```
### go build
```
go build github.com/wybiral/hookah/cmd/hookah
```

# Usage instructions (CLI)
The hookah command allows you to pipe data between various sources/destinations.
By default pipes are full duplex but can be limited to input/output-only mode.

For details run `hookah -h`

## Examples

Pipe from stdin/stdout to a new TCP listener on port 8080:
```
hookah stdio tcp-listen://localhost:8080
```
Note: this is the same even if you ommit the `stdio` part because hookah will
assume stdio is indended when only one node (tcp-listen in this case) is used.

Pipe from a TCP client on port 8080 to a new WebSocket listener on
port 8081:
```
hookah tcp://localhost:8080 ws-listen://localhost:8081
```

Pipe from a new Unix domain socket listener to a TCP client on port 8080:
```
hookah unix-listen://path/to/sock tcp://localhost:8080
```

Pipe only the input from a TCP client to the output of another TCP client:
```
hookah -i tcp://:8080 -o tcp://:8081
```

Fan-out the input from a TCP listener to the output of multiple TCP clients:
```
hookah -i tcp-listen://:8080 -o tcp://:8081 -o tcp://:8082 -o tcp://:8083
```

# Usage instructions (Go package)
[See godoc page](https://godoc.org/github.com/wybiral/hookah).
