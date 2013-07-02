format:
	gofmt -w .
	git add -A
	git commit -m 'gofmt'
	git push origin HEAD
get-deps:
	go get github.com/bmizerany/pq
	go get github.com/alecthomas/gozmq
	go get github.com/gorilla/context
	go get github.com/gorilla/mux
	go get github.com/gorilla/rpc
	go get code.google.com/p/go.net
	go get code.google.com/p/goprotobuf
sloc:
	cloc . | grep Go
