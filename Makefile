all: iterchange

bin/staticcheck: go.mod go.sum
	GOBIN=`pwd`/bin go install honnef.co/go/tools/cmd/staticcheck

bin: bin/iterchange 

iterchange:
	go build .

clean-bin:
	rm -rf bin
