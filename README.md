## gopoker

### Features

#### Game variations
* [**Holdem poker**](http://en.wikipedia.org/wiki/Community_card_poker): Texas, Omaha, Omaha Hi/Lo
* [**7-card poker**](http://en.wikipedia.org/wiki/Stud_poker): Stud, Stud Hi/Lo, Razz, London lowball
* [**Draw poker**](http://en.wikipedia.org/wiki/Draw_poker): 5-card draw, 2-7 single and triple draw, [Badugi](http://en.wikipedia.org/wiki/Badugi)
* **Mixed games**: [H.O.R.S.E.](en.wikipedia.org/wiki/HORSE), 8-game

#### Limits
* Fixed Limit
* Pot Limit
* No Limit

#### HTTP server
default port is 8080

* `/_api` - REST API
* `/_ws` - websockets
* `/_rpc` - JSON RPC

Sample node config:

```json
{
    "HTTP": {
        "Addr": ":8080",
        "APIPath": "/_api",
        "RPCPath": "/_rpc",
        "WebSocketPath": "/_ws"
    },
    "RPC": {
        "Addr": ":8081",
        "Timeout": 5
    },
    "Logdir": "/var/log/gopoker"
}
```

#### IPC
* pubsub via 0mq sockets using binary protocol (Google protobuf), default port is 5555
* JSON RPC via TCP socket, default port is 8081

#### Tools
* `gopoker-cli` - REPL-style gameplay
* `gopoker-ctrl` - command line interface to RPC service
* `gopoker-bot` - configurable bot with simple AI
* web client - see https://github.com/maliqq/poker-js

### Architecture
* `ai/` - bot AI with decision making logic
* `calc/` - poker related math
* `client/` - client related code
* `model/` - poker domain
* `play/` - gameplay
* `poker/` - poker rules
* `protocol/` - poker events exchange
* `protocol/message` - poker protocol
* `server/` - server with topology (node, cluster, balancer) and services
* `storage/` - persistence to PostgreSQL for critical data and MongoDB for temporary data
* `util/` - utility functions
