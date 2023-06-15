# dockerfiles

build contextがリポジトリのルートディレクトリの前提で書く

## prism.Dockerfile

[Prism](https://github.com/stoplightio/prism/tree/master)でモックサーバーを立てるためのイメージ
`docs/openapi.yaml`を参照して起動する

### ビルドとコンテナの起動
```bash
# リポジトリのルートディレクトリで実行
docker build -f dockerfiles/prism.Dockerfile -t h23s_15-prism .
docker run -p 4010:4010 h23s_15-prism mock -h 0.0.0.0 /tmp/openapi.yaml
```

`http://localhost:4010`でアクセス可能