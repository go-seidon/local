# local-storage

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=go-seidon_local&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=go-seidon_local)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=go-seidon_local&metric=coverage)](https://sonarcloud.io/summary/new_code?id=go-seidon_local)

## Doc
No doc right now

## Todo
1. Allow injecting clock in repository
2. Allow debugging using `APP_DEBUG` = true

## Nice to have
1. Separate findFile query in DeleteFile and RetrieveFile
2. File meta for storing file related data, e.g: user_id, feature, category, etc

## Technical Stack
1. Transport layer
- rest
- grpc
2. Database
- mysql
- postgres
- mongo
3. Config
- system environment

## Run
### Docker
1. First running: `docker-compose up -d`
2. MySQL database: 
- `docker-compose up mysql-database`
- `docker-compose stop mysql-database`
3. MySQL Test database:
- `docker-compose up mysql-test-database`
- `docker-compose stop mysql-test-database`
