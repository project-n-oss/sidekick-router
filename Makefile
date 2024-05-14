all: iterchange

bin/staticcheck: go.mod go.sum
	GOBIN=`pwd`/bin go install honnef.co/go/tools/cmd/staticcheck@latest

bin: bin/staticcheck 

iterchange:
	go build .

clean-bin:
	rm -rf bin
