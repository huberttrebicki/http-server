# http-server

A minimal HTTP/1.1 server built from scratch over TCP.
It has a request parsing based on the RFC guidelines and a basic routing


## Usage
Start the server:

```sh
go run ./cmd/httpserver
```

This listens on port `42069`. You can test it with curl:

```sh
curl http://localhost:42069/
```
