FROM golang:1.21-alpine

WORKDIR /go/src/github.com/project-n-oss/sidekick-router
COPY . .

RUN go generate ./...
RUN go build .

FROM golang:1.21-alpine

WORKDIR /usr/bin

COPY --from=0 /go/src/github.com/project-n-oss/sidekick-router/sidekick-router .
RUN ./sidekick-router --help > /dev/null

ENTRYPOINT ["/usr/bin/sidekick-router"]

