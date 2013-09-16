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

build-all:
	protoc --go_out=. exch/message/*.proto
	go build gopoker/bin/gopoker-bot
	go build gopoker/bin/gopoker-cli
	go build gopoker/bin/gopoker-ctrl
	go build gopoker/bin/gopoker-server
	gofmt -w .

test-all:
	go test gopoker/ai
	go test gopoker/model
	go test gopoker/play
	go test gopoker/play/context
	go test gopoker/exch/message

install-all:
	go install gopoker/bin/gopoker-bot
	go install gopoker/bin/gopoker-cli
	go install gopoker/bin/gopoker-ctrl
	go install gopoker/bin/gopoker-server

clean-all:
	rm exch/message/*.pb.go
	rm gopoker-*
