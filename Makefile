get-deps:
	go get github.com/bmizerany/pq
	go get github.com/gorilla/context
	go get github.com/gorilla/mux
	go get github.com/gorilla/rpc
	go get code.google.com/p/go.net/websocket
	go get code.google.com/p/goprotobuf/proto
	go get code.google.com/p/goprotobuf/protoc-gen-go
	go get -tags zmq_3_x github.com/alecthomas/gozmq
	go get labix.org/v2/mgo

build-all:
	protoc --go_out=. protocol/message/*.proto
	go build gopoker/bin/gopoker-cli
	go build gopoker/bin/gopoker-ctrl
	go build gopoker/bin/gopoker-server
	gofmt -w .

test-all:
	go test gopoker/model
	go test gopoker/play
	go test gopoker/play/context

install-all:
	go install gopoker/bin/gopoker-cli
	go install gopoker/bin/gopoker-ctrl
	go install gopoker/bin/gopoker-server
