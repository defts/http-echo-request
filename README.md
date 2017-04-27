# http-echo-request

Will accept any TCP connection and echo back a HTTP response with the
entire content of the incoming TCP connection.

## Run

```
PORT=9090 go run http-echo-request.go
```

## Example usage

Just curl:

```
curl http://127.0.0.1
```

## License

MIT
