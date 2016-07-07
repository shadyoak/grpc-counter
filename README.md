# gRPC Counter

A simple counter demonstration using gRPC.

http://www.grpc.io

## Installation

```console
$ go get -u github.com/shadyoak/grpc-counter
```

## Running the Demo

First, start the grpc-counter server.

```console
$ grpc-counter
```

Compile the demo client.

```console
$ cd demo/
$ go build
```

Open your web browser to `http://localhost:8080` and run the demo.

```console
$ ./demo
```

## Extra Credit

Run multiple instances of the demo simultaneously. What happens?
