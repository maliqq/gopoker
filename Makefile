get-deps:
	go get github.com/bmizerany/pq
	go get github.com/alecthomas/gozmq
	go get github.com/gorilla/context
	go get github.com/gorilla/mux
	go get github.com/gorilla/rpc
	go get code.google.com/p/go.net
	#go get code.google.com/p/goprotobuf
	go get code.google.com/p/goprotobuf/{proto,protoc-gen-go}
	go get -tags zmq_3_x github.com/alecthomas/gozmq

build-all:
	go build gopoker/gopoker-console
	go build gopoker/gopoker-control
	go build gopoker/gopoker-server
	protoc --go_out=. *.proto

test-all:
	go test gopoker/model
	go test gopoker/play
	go test gopoker/play/context

install-all:
	go install gopoker/bin/gopoker-console
	go install gopoker/bin/gopoker-control
	go install gopoker/bin/gopoker-server
