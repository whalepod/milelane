[![codecov](https://codecov.io/gh/whalepod/milelane/branch/master/graph/badge.svg)](https://codecov.io/gh/whalepod/milelane)

# milelane

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

- Healthcheck
```
http://docker-local.milelane.co/tasks
# => Right after init, it returns null.
```

## Setup Database
```
$ docker exec -i -t milelane_mysql_1 bash
$ mysql -u root -h 127.0.0.1 milelane < docker-entrypoint-initdb.d/schema.sql
```

## How to test, lint and auto formatting.

### test
- `make test`
    - This command also exports coverage report. To check coverage, `open cover.html`

### auto formatting
- `make fmt`

### check lint
- `make lint`

### vet
- `make vet`

## How to migrate
We use [goose](https://github.com/pressly/goose) to migrate.

### Set env variable
You need to set environment variable.

#### dev
export USER_NAME=root
export DATABASE=milelane

### Command
#### Create migration file
- `make migrate-create NAME=xxx_yyy_zzz`

#### Apply migration
- `migrate-up`
