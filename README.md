## gopoker

### Features

#### Games
* Holdem poker: Texas, Omaha, Omaha Hi/Lo
* 7-card poker: Stud, Stud Hi/Lo, Razz, London lowball
* Draw poker: 5-card draw, 2-7 single and triple draw, Badugi
* Mixed games: H.O.R.S.E., 8-game

#### Limits
* Fixed Limit
* Pot Limit
* No Limit


#### HTTP server
default port is 8080

* /_api - REST API 
* /_ws - websockets
* /_rpc - JSON RPC

#### IPC
* pubsub via 0mq sockets using binary protocol (Google protobuf), default port is 5555
* JSON RPC via TCP socket, default port is 8081

#### Tools
* gopoker-cli - REPL-style gameplay
* gopoker-ctrl - command line interface to RPC service
* gopoker-bot - configurable bot with simple AI
* web client - see https://github.com/maliqq/poker-js
