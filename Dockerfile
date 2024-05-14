FROM golang:1.20-alpine

WORKDIR /go/src/github.com/project-n-oss/interchange
COPY . .

RUN go generate ./...
RUN go build .

FROM golang:1.20-alpine

WORKDIR /usr/bin

COPY --from=0 /go/src/github.com/project-n-oss/interchange/interchange .
RUN ./interchange --help > /dev/null

ENTRYPOINT ["/usr/bin/interchange"]

