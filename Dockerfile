# Client build
FROM node:24.16.0-alpine@sha256:21f403ab171f2dc89bad4dd69d7721bfd15f084ccb46cdd225f31f2bc59b5c9a AS client-build

WORKDIR /app

COPY client/package.json client/pnpm-lock.yaml client/pnpm-workspace.yaml ./
RUN corepack enable \
    && pnpm install --frozen-lockfile

COPY client/ .
RUN pnpm run build


# Server build
FROM golang:1.25.0-alpine@sha256:f18a072054848d87a8077455f0ac8a25886f2397f88bfdd222d6fafbb5bba440 AS server-build

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
