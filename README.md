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

`docker compose up --watch`でアプリケーションとMariaDBを起動します。
サーバーとクライアントのファイル変更時はアプリケーションイメージを再ビルドします。
OpenAPI UIは通常の開発環境では起動しません。

#### OpenAPI UIの立ち上げ
```
task openapi-ui
```

`http://localhost:8100`で`docs/openapi.yaml`を確認できます。

#### openapi-codegenの実行
```
task install-server-openapi-codegen
task server-openapi-codegen
```

### クライアントサイド
