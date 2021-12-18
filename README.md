# Golang で ローカル Azure Storage (Azurite) にアクセスする

## Reference

* [azurite](https://docs.microsoft.com/ja-jp/azure/storage/common/storage-use-azurite?tabs=docker-hub#authorization-for-tools-and-sdks)

## 概要

本リポジトリは、Golang で Azure Storage のエミュレーターである Azurite にアクセスすることを可能にする。

|Storage|Status of implementation|
|:--|:--|
|Azure Blob||
|Queue Storage|✓|
|Table Storage||

## Azure Queue Storage へのアクセス

### 構成

作業ディレクトリ構成は以下の通りである。

```
.
└── azurite-controller/
     ├── certs/
     │    ├── 127.0.0.1-key.pem
     │    └── 127.0.0.1.pem
     ├── docker-compose.yml
     └── queue_test.go
```

#### 本リポジトリのクローン

1. 本リポジトリを作業ディレクトリにクローンする。

#### HTTPS 設定

1. [mkcert](https://github.com/FiloSottile/mkcert) を利用した以下のスクリプトを実行し、自己署名証明書を作成する。
   ```sh
   $ chmod 755 certs/create.sh
   $ ./cert/create.sh
   ```

#### Docker Compose 実行

1. 以下のコマンドを実行し、Azurite コンテナを起動する。
   ```sh
   $ docker-compose up -d
   ```

#### テストコード実行

1. 以下のコマンドを実行し、テストを実行する。
   ```sh
   $ go test -v azuritectl_test.go
   ```

