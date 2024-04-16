# go-stream-idle-alert

This is a simple utility that will alert you when a stream goes idle. It is useful for monitoring a stream that should be active at all times. 

The monitor stream can be updated by the settings.

## Start

Run the main file:

```bash
go run cmd/app/main.go
```

## Build

Build the binary:

```bash
go build -o go-stream-idle-alert cmd/app/main.go
```

Run the binary:

```bash
./go-stream-idle-alert
```