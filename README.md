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
├── azurite-controller/
│    ├── azurite/
│    │    ├── 127.0.0.1-key.pem
│    │    └── 127.0.0.1.pem
│    ├── queue/
│    ├── docker-compose.yml
│    └── queue_test.go
└── mkcert/
```

#### 本リポジトリのクローン

1. 本リポジトリを作業ディレクトリにクローンする。

#### HTTPS 設定

1. [mkcert installation](https://github.com/FiloSottile/mkcert#installation) を参考に以下の手順を実行する。

   ```sh
   $ git clone https://github.com/FiloSottile/mkcert && cd mkcert

   $ go build -ldflags "-X main.Version=$(git describe --tags)"

   $ ./mkcert -install

   $ ./mkcert 127.0.0.1
     => 127.0.0.1.pem and 127.0.0.1-key.pem created.

   $ mv 127.0.0.1.pem     ../azurite-controller/azurite/127.0.0.1.pem

   $ mv 127.0.0.1-key.pem ../azurite-controller/azurite/127.0.0.1-key.pem
   ```

#### Docker Compose 実行

1. 以下のコマンドを実行する。

   ```sh
   $ docker-compose up -d
   ```

#### テストコード実行

1. 以下のコマンドを実行する。

   ```sh
   $ go test -v queue_test.go
   ```

