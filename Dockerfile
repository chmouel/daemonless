FROM openshift/origin-release:golang-1.13 AS builder
RUN
WORKDIR /go/src/github.com/chmouel/daemonless
COPY . .
RUN go build ./...
