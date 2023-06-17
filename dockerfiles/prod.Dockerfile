# クライアントサイドのビルド
FROM node:18.16-alpine3.18 as client-build

WORKDIR /app

COPY client/package*.json .
RUN npm ci

COPY client/ .

RUN npm run build


# サーバーサイドのビルド
FROM golang:1.20.5-alpine3.18 as server-build

EXPOSE 8080
EXPOSE 8100

WORKDIR /github.com/traP-jp/h23s_15

COPY server/go.mod .
COPY server/go.sum .
RUN go mod download

COPY server/ .

RUN go build -o app .


# 最終的な配信用
FROM caddy:2.6.4-alpine

WORKDIR /

COPY config/Caddyfile /etc/caddy/Caddyfile

COPY --from=client-build /app/dist /usr/share/caddy
COPY --from=server-build /github.com/traP-jp/h23s_15/app app

EXPOSE 80

ENTRYPOINT [ "./app" ]