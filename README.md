# HTTP Server in Go

A small HTTP/1.1 server built from scratch on top of TCP, without using Go's `net/http` server implementation.

## Features

- HTTP/1.1 request-line parsing and validation.
- Incremental parsing across fragmented network reads.
- Case-insensitive header storage and lookup.
- Validation of HTTP header names.
- Support for duplicate header fields.
- Request-body parsing using `Content-Length`.
- Response writing for status lines, headers, and bodies.
- `200 OK`, `400 Bad Request`, and `500 Internal Server Error` responses.
- Callback-based request handlers.
- Concurrent connections using one goroutine per client.
- Automatic `400 Bad Request` responses for malformed requests.

## Project structure

| Path | Responsibility |
| --- | --- |
| `cmd/httpserver` | Example HTTP server and routes. |
| `cmd/tcplistener` | Low-level TCP request parsing example. |
| `internal/server` | TCP listener and connection handling. |
| `internal/request` | Request-line, header, and body parsing. |
| `internal/headers` | Header validation, normalization, and storage. |
| `internal/response` | HTTP response serialization. |

## Running

The example server requires Go 1.26.1 or newer.

```sh
go run ./cmd/httpserver
```

It listens on port `42069`. Try the included routes with `curl`:

```sh
curl -i http://localhost:42069/
curl -i http://localhost:42069/yourproblem
curl -i http://localhost:42069/myproblem
```

These return `200 OK`, `400 Bad Request`, and `500 Internal Server Error`, respectively.

To build the executable:

```sh
go build -o httpserver ./cmd/httpserver
./httpserver
```

## Testing

Run the complete test suite with:

```sh
go test ./...
```

The tests cover request lines, HTTP headers, request bodies, malformed input, and parsing data delivered in small chunks.
