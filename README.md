# h23s_15

## 開発用

### 共通

#### Taskのインストール
必須です

```
go install github.com/go-task/task/v3/cmd/task@v3.26.0
```

### サーバーサイド

開発環境の立ち上げ
```
task server-dev
```

openapi-codegenの実行
```
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.11.0
oapi-codegen -package api docs/openapi.yaml > server/api/server.gen.go
```

### クライアントサイド