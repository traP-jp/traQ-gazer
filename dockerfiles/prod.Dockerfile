FROM golang:1.20.5-alpine3.18

EXPOSE 8080
EXPOSE 8100

WORKDIR /github.com/traP-jp/h23s_15

COPY server/go.mod .
COPY server/go.sum .
RUN go mod download

COPY server/ .

RUN go build -o app .

CMD [ "app" ]

# TODO: マルチステージビルドでクライアントもビルドする