# Golang Socket.IO

This library was forked from github.com/graarh/golang-socketio

It provides a simple Golang implementation of
[socket.io](http://socket.io) client and server.

The code was heavily refactored, also implemented XHR polling transport
for client and server upgrade XHR -> websocket.

**Pull requests appreciated**

## Usage examples

Please observe the `examples` directory for usage examples:

```
JavaScript client:    examples/assets/index.html, serve it with:
Go server:            go run examples/server/server.go

Go client via WS:     go run examples/client_websocket/client.go
Go client via XHR:    go run examples/client_xhr_polling/client.go
```

Please note that no Go client upgrade implemented yet.

This client is mainly for testing purposes.

## Installation

    go get github.com/mtfelian/golang-socketio

## TODOs, ideas to further development

- write tests, make a good test coverage
- Go client's upgrade from XHR to WS
- Go server's ability to fallback from WS to XHR
- Go client's ability to fallback from WS to XHR
- support newer versions of socket.io protocol
