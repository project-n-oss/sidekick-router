all: sidekick-router

bin/staticcheck: go.mod go.sum
	GOBIN=`pwd`/bin go install honnef.co/go/tools/cmd/staticcheck@latest

bin: bin/staticcheck 

.PHONY: sidekick-router clean-bin
sidekick-router:
	go build .

clean-bin:
	rm -rf bin
