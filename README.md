[![codecov](https://codecov.io/gh/whalepod/milelane/branch/master/graph/badge.svg)](https://codecov.io/gh/whalepod/milelane)

# milelane

## How to test

- `go test -coverprofile=cover.out ./... && go tool cover -html=cover.out -o cover.html`

## Setup with docker-compose

- Install Docker
- Update /etc/hosts, add setting below.
```
127.0.0.1 docker-local.milelane.co
```

- Run
```
docker-compose up -d
```

- Access
```
http://docker-local.milelane.co/tasks
# => Right after init, it returns null.
```

## Setup Database
```
$ docker exec -i -t milelane_mysql_1 bash
$ mysql -u root -h 127.0.0.1 milelane < docker-entrypoint-initdb.d/schema.sql
```
