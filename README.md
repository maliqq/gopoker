## gopoker
>Golang poker server 

>http://pokerno.de/

### Features

#### Game variations
* [**Holdem poker**](http://en.wikipedia.org/wiki/Community_card_poker): Texas, Omaha, Omaha Hi/Lo
* [**7-card poker**](http://en.wikipedia.org/wiki/Stud_poker): Stud, Stud Hi/Lo, Razz, London lowball
* [**Draw poker**](http://en.wikipedia.org/wiki/Draw_poker): 5-card draw, 2-7 single and triple draw, [Badugi](http://en.wikipedia.org/wiki/Badugi)
* **Mixed games**: [H.O.R.S.E.](http://en.wikipedia.org/wiki/HORSE), 8-game

#### Limits
* Fixed Limit
* Pot Limit
* No Limit

#### HTTP server
default port is 8080

* `/_api` - REST API
* `/_ws` - websockets
* `/_rpc` - JSON RPC

Sample node config ([`node.json`](etc/node.json)):

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

#### Protocols
* pubsub service via 0mq socket+Google protobuf, default port is 5555
* JSON RPC service, default port is 8081
* TODO: STOMP
* TODO: BOSH

#### Tools
* `bin/gopoker-cli` - REPL-style gameplay
* `bin/gopoker-ctrl` - command line interface to RPC service
* `bin/gopoker-bot` - configurable bot with simple AI
* web client - see https://github.com/maliqq/poker-js

### Architecture
* [`ai/`](ai/) - bot AI with decision making logic
* [`calc/`](calc/) - poker related math
* [`client/`](client/) - client related code
* [`exch/`](exch/) - events exchange
 * [`exch/message/`](exch/message/) - message format
 * [`exch/protocol`](exch/protocol/) - protocols
* [`model/`](model/) - poker domain model
* [`play/`](play/) - gameplay
* [`poker/`](poker/) - poker rules
* [`server/`](server/) - server with topology (node, cluster, balancer) and services
* [`storage/`](storage/) - persistence to PostgreSQL for critical data and MongoDB for temporary data
* [`util/`](util/) - utility functions
