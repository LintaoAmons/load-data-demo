# 造 1000000 数据并载入到数据库，总耗时2分48s

这里的样例数据结构比较简单，但是在我的真实项目中，造 几百万条数据并加载进远端数据库也仅耗时 10 分钟以内

```
go--load-data-to-db on  go--load-data-to-db [!?] via 🐳 lima-docker_rootful_x86 via
 🐹 v1.20.6 
❯ ./load.sh
COPY 100000
COPY 100000
COPY 100000
COPY 100000
COPY 100000
COPY 100000
COPY 100000
COPY 100000
COPY 100000
COPY 100000

go--load-data-to-db on  go--load-data-to-db [!?] via 🐳 lima-docker_rootful_x86 via
 🐹 v1.20.6 took 2m48s 
```


## generate struct out of db schema by jet
> https://github.com/go-jet/jet

0. make sure db is running, mock this by docker-compose

1. install jet

```
go install github.com/go-jet/jet/v2/cmd/jet@latest
```

2. generate

```
jet -source=postgres -host=localhost -port=5499 -user=world -password=world123 -dbname=world-db -path=./gen
```

3. run main file

## load csv data into db

```
PGPASSWORD=world123 psql -h localhost -p 5499 -d world-db -U world -c "\copy public.city from 'city-1.csv' with (format csv);"
```

