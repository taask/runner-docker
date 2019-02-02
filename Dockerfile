FROM golang:1.11.2 as builder

RUN mkdir -p /go/src/github.com/taask/runner-docker
WORKDIR /go/src/github.com/taask/runner-docker

COPY . .

RUN CGO_ENABLED=0 go build

FROM taask/docker-sandbox:latest

RUN mkdir -p /taask/runner/data

COPY --from=builder /go/src/github.com/taask/runner-docker/runner-docker /taask/runner-docker