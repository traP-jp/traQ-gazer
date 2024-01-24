# traQ-gazer

## 開発用

### 共通

#### Taskのインストール
必須です

```
go install github.com/go-task/task/v3/cmd/task@v3.26.0
```

### サーバーサイド

#### 開発環境の立ち上げ
```
task server-dev
```

#### openapi-codegenの実行
```
task install-server-openapi-codegen
task server-openapi-codegen
```

### クライアントサイド