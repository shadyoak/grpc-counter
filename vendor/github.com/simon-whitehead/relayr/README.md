## RelayR

RelayR is a Go package that provides easy-to-use real time communication APIs for Go web applications.

#### Current version: 0.3.0

Please see CHANGELOG.md for details about the changes between versions.

### Installation

#### IMPORTANT:

RelayR relies on the [Gorilla WebSocket package](https://github.com/gorilla/websocket) for WebSocket support. Please head over to that repository, give it a star, and ``go get` it into your `GOPATH`.

RelayR itself however, will fall back to Long Polling for any browsers that do not support Web Sockets (you just need the above package so that your server supports Web Sockets).

#### Installing RelayR

After you have installed the [Gorilla WebSocket package](https://github.com/gorilla/websocket), you can run the following:

    go get -u github.com/simon-whitehead/relayr

### Examples

The `/examples` directory contains three examples.

The first one is a simple Server -> Client timestamp push. The server will push the current time down to all connected clients every second.

It looks like this:

![relayr-time-push](https://cloud.githubusercontent.com/assets/2499070/6539845/2a2d7d0a-c4d5-11e4-8ace-9f619769dca9.gif)

The second one is a replica of the famous SignalR sample, "High-Frequency Realtime with SignalR" where a shape can be dragged around a browser and broadcast to all clients.

It looks like this:

![relayr-shape-move](https://cloud.githubusercontent.com/assets/2499070/6540051/ac091b60-c4dd-11e4-9115-9debcd836136.gif)

The third one demonstrates Groups and membership. It allows clients to subscribe to groups and receive notifications that are pushed to groups they are subscribed to.

### Project status

This is very much a work in progress - some might say, pre-alpha. The code is horrible .. but it works. I am opening it up to the world early for feedback and contributions (in any form).

### Contributing

Please! Fork away .. improve it .. create a PR .. whatever you want to do. I am open to any and all help.

### Licence

This code is released under the Apache 2.0 licence.
