# hookah
A Swiss Army knife for data pipelines.

[![GoDoc](https://godoc.org/github.com/wybiral/hookah?status.svg)](https://godoc.org/github.com/wybiral/hookah)

### Demo video
[![Watch demo](https://img.youtube.com/vi/AdFzK9sDywg/0.jpg)](https://www.youtube.com/watch?v=AdFzK9sDywg)

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
The hookah command allows you to specify an input source -i and an output destination -o.
Any data that's fed into the input will be piped to the output.

## Examples

Pipe from stdin to a new TCP listener on port 8080:
```
hookah -o tcp-listen://localhost:8080
```

Pipe from an existing TCP listener on port 8080 to a new HTTP listener on port 8081:
```
hookah -i tcp://localhost:8080 -o http-listen://localhost:8081
```

Pipe from a new Unix domain socket listener to stdout:
```
hookah -i unix-listen://path/to/sock
```

Pipe from a new HTTP listener on port 8080 to an existing Unix domain socket:
```
hookah -i http-listen://localhost:8080 -o unix://path/to/sock
```

# Usage instructions (Go package)
[See godoc page](https://godoc.org/github.com/wybiral/hookah).
