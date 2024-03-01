# Backend application IDLER chat service

###
- GO 1.18.2
- GIN
- PGX
- GCACHE

For application need EnvFile by Borys Pierov plugin and .env file which contains:
```dotenv
HTTP_HOST=host.docker.internal
HTTP_PORT=[your application port here]

POSTGRES_VERSION=14
POSTGRES_PORT=[your postgres port here]
POSTGRES_DB=idler
POSTGRES_SCHEMA=idler
POSTGRES_URL=jdbc:postgresql://${HTTP_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?currentSchema=${POSTGRES_SCHEMA}
POSTGRES_USER=[your postgres user here]
POSTGRES_PASSWORD=[your postgres password here]

LIQUIBASE_VERSION=4.11

GRAFANA_VERSION=9.0.2
GRAFANA_USER=[your grafana user here]
GRAFANA_PASSWORD=[your grafana password here]
GRAFANA_PORT=[your grafana port here]

PROMETHEUS_VERSION=v2.36.2
PROMETHEUS_PORT=[your prometheus port here]

WEBSOCKET_PORT=[your websocket port here]

MONGO_VERSION=4.4.6
MONGO_HOST=[your mongo host here]
MONGO_PORT=[your mongo port here]
MONGO_DB=[your mongo db here]
MONGO_INITDB_ROOT_USERNAME=[your mongo username here]
MONGO_INITDB_ROOT_PASSWORD=[your mongo password here]

HASH_SALT=[your salt here]
JWT_SIGNING_KEY=[your signing key here]
```
For successfully running liquibase need to append in db/liquibase.properties:
```dotenv
username: [your postgres user here]
password: [your postgres password here]
```
Command for building application
```dotenv
- make build
```
Command for running tests application
```dotenv
- make build
```

Command for running docker containers
```dotenv
- make docker
```

Swagger documentation http://localhost:9000/swagger/index.html
