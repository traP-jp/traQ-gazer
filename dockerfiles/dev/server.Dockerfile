FROM golang:1.25.0@sha256:5502b0e56fca23feba76dbc5387ba59c593c02ccc2f0f7355871ea9a0852cebe

WORKDIR /github.com/traP-jp/h23s_15

RUN apt install -y git \
    && go install github.com/cosmtrek/air@v1.44.0

COPY server/go.mod .
COPY server/go.sum .
RUN go mod download
