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
### dev
export USER_NAME=root
export DATABASE=milelane
