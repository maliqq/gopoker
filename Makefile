get-deps:
	go get github.com/bmizerany/pq
	go get github.com/gorilla/context
	go get github.com/gorilla/mux
	go get github.com/gorilla/rpc
	go get code.google.com/p/go.net/websocket
	go get code.google.com/p/goprotobuf/proto
	go get code.google.com/p/goprotobuf/protoc-gen-go
	go get -tags zmq_2_x github.com/alecthomas/gozmq
	go get labix.org/v2/mgo
	go get github.com/hoisie/redis
	go get code.google.com/p/go.crypto/bcrypt
	go get github.com/rcrowley/go-metrics
	go get github.com/jjeffery/stomp
	go get github.com/streadway/amqp
	go get github.com/vmihailenco/msgpack

build-all:
	#protoc --go_out=. event/message/protobuf/*.proto
	go build gopoker/bin/gopoker-bot
	go build gopoker/bin/gopoker-cli
	go build gopoker/bin/gopoker-ctrl
	go build gopoker/bin/gopoker-server
	gofmt -w .

test-all:
	go test gopoker/test/ai
	go test gopoker/test/model
	go test gopoker/test/play
	go test gopoker/test/play/context
	go test gopoker/test/event/message

install-all:
	go install gopoker/bin/gopoker-bot
	go install gopoker/bin/gopoker-cli
	go install gopoker/bin/gopoker-ctrl
	go install gopoker/bin/gopoker-server

clean-all:
	rm event/message/protobuf/*.pb.go
	rm gopoker-*
