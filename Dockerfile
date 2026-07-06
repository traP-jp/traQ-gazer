# Client build
FROM node:26.4.0-alpine@sha256:725aeba2364a9b16beae49e180d83bd597dbd0b15c47f1f28875c290bfd255b9 AS client-build

WORKDIR /app

RUN npm install -g corepack@0.35.0 \
    && corepack enable

RUN --mount=type=bind,source=client/package.json,target=/app/package.json,readonly \
    --mount=type=bind,source=client/pnpm-lock.yaml,target=/app/pnpm-lock.yaml,readonly \
    --mount=type=bind,source=client/pnpm-workspace.yaml,target=/app/pnpm-workspace.yaml,readonly \
    --mount=type=cache,id=traq-gazer-pnpm-store,target=/pnpm/store,sharing=locked \
    pnpm fetch --frozen-lockfile --store-dir=/pnpm/store

RUN --mount=type=bind,source=client,target=/app,rw \
    --mount=type=cache,id=traq-gazer-pnpm-store,target=/pnpm/store,sharing=locked \
    mkdir -p /out \
    && pnpm install --frozen-lockfile --prefer-offline --store-dir=/pnpm/store \
    && pnpm build --outDir /out/dist --emptyOutDir


# Server build
FROM golang:1.26.4-alpine@sha256:3ad57304ad93bbec8548a0437ad9e06a455660655d9af011d58b993f6f615648 AS server-build

WORKDIR /app

RUN --mount=type=bind,source=server/go.mod,target=/app/go.mod,readonly \
    --mount=type=bind,source=server/go.sum,target=/app/go.sum,readonly \
    --mount=type=cache,target=/go/pkg/mod \
    go mod download

RUN --mount=type=bind,source=server,target=/app,readonly \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    mkdir -p /out \
    && go build -mod=readonly -buildvcs=false -o /out/app .


# Production app image
FROM alpine:3.24.1@sha256:28bd5fe8b56d1bd048e5babf5b10710ebe0bae67db86916198a6eec434943f8b AS app

WORKDIR /

COPY --from=client-build /out/dist dist
COPY --from=server-build /out/app app

EXPOSE 8080

ENTRYPOINT ["./app"]
