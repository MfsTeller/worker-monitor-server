# Postgresql

## コンテナ操作方法

- コンテナ起動

~~~
# docker-compose.exe up -d
~~~

- コンテナにログイン

~~~
# winpty docker exec -ti postgres //bin/bash
~~~

- コンテナ削除

~~~
# docker-compose.exe down
~~~

## DB操作方法

- DBにログイン

~~~
# psql -h <host> -p <port> -U <user> -d <dbname>

# Example
# psql -h localhost -p 5432 -U root -d workermonitordb
~~~