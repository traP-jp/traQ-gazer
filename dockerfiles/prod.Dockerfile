# クライアントサイドのビルド
FROM node:21.7-alpine3.18 as client-build

WORKDIR /app

COPY client/package*.json .
RUN npm ci

COPY client/ .

RUN npm run build


# サーバーサイドのビルド
FROM golang:1.20.5-alpine3.18 as server-build

WORKDIR /github.com/traP-jp/h23s_15

COPY server/go.mod .
COPY server/go.sum .
RUN go mod download

COPY server/ .

RUN go build -o app .


# 最終的な配信用
FROM alpine:3.19.1

WORKDIR /

COPY --from=client-build /app/dist dist
COPY --from=server-build /github.com/traP-jp/h23s_15/app app

EXPOSE 8080
EXPOSE 8100

ENTRYPOINT [ "./app" ]