# Client build
FROM node:26.4.0-alpine@sha256:725aeba2364a9b16beae49e180d83bd597dbd0b15c47f1f28875c290bfd255b9 AS client-build

WORKDIR /app

COPY client/package.json client/pnpm-lock.yaml client/pnpm-workspace.yaml ./
RUN npm install -g corepack@0.35.0 \
    && corepack enable \
    && pnpm install --frozen-lockfile

COPY client/ .
RUN pnpm run build


# Server build
FROM golang:1.26.4-alpine@sha256:3ad57304ad93bbec8548a0437ad9e06a455660655d9af011d58b993f6f615648 AS server-build

WORKDIR /github.com/traP-jp/h23s_15

COPY server/go.mod server/go.sum ./
RUN go mod download

COPY server/ .
RUN go build -buildvcs=false -o /app/app .


# Production app image
FROM alpine:3.24.1@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b AS app

WORKDIR /

COPY --from=client-build /app/dist dist
COPY --from=server-build /app/app app

EXPOSE 8080

ENTRYPOINT ["./app"]
