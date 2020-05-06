# Worker-Monitor Server

本ソフトウェアはWorker-Monitor Clientのデータを蓄積・管理します。

## Background・Purpose

[worker-monitor-clientリポジトリ](https://github.com/MfsTeller/worker-monitor-client)のBackgroundを参照してください。

## Getting started

### 1. DBコンテナのデプロイ

下記のコマンドを実行すると`postgres`コンテナがデプロイされます。

~~~
# cd postgresql
# docker-compose.exe up -d
~~~

DBの状態を確認したい時などは、下記のコマンドを実行し、コンテナにログインしてください。

~~~
# cd postgresql
# winpty docker exec -ti postgres //bin/bash
~~~

DBにログインする場合は下記のコマンドを実行してください。

~~~
# Format
# psql -h <host> -p <port> -U <user> -d <dbname>

# Example
# psql -h localhost -p 5432 -U root -d workermonitordb
~~~

コンテナを削除するときは下記のコマンドを実行してください。

~~~
# cd postgresql
# docker-compose.exe down
~~~

## 2. Worker-Monitor Serverコンテナのデプロイ

下記のコマンドを実行し、コンテナイメージをビルドしてください。

~~~
# docker build -t openapi-server .
~~~

下記のコマンドを実行すると、コンテナが立ち上がります。

~~~
# docker run -p 8080:8080 openapi-server
~~~

APIに修正が必要な場合など、`openapi.yaml`を修正した場合は、下記のコマンドを実行し、ソースコードを生成してください。（[参考： OpenAPI Generator - Docker](https://github.com/OpenAPITools/openapi-generator#16---docker)）

~~~
# docker run --rm -v /$PWD:/local openapitools/openapi-generator-cli generate -i local/openapi.yaml -g go-server -o local/go-server
~~~

### Worker-Monitor ClientData API

- `GET /clientdata/{clientId}`

~~~
$ curl http://192.168.99.100:8080/clientdata/2

[
    {
        client_id:2,
        name:Taro Ziro,
        startup_datetime:2020/04/30 00:00:00,
        shutdown_datetime:2020/04/30 08:00:00
    }
]
~~~

- `POST /clientdata`

~~~
$ curl -X POST -H "Content-Type: application/json" \
  -d "[{\"client_id\":2,\"name\":\"Taro Ziro\",\"startup_datetime\":\"2020/04/30 00:00:00\",\"shutdown_datetime\":\"2020/04/30 08:00:00\"}]" \
  http://192.168.99.100:8080/clientdata
~~~

