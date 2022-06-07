# local-storage

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=go-seidon_local&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=go-seidon_local)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=go-seidon_local&metric=coverage)](https://sonarcloud.io/summary/new_code?id=go-seidon_local)

## Doc
No doc right now

## Todo
Nothing todo right now

## Business Flow
1. Client can upload a file 
- required: id
- optional: name (default: originame name), visibility(public, private, default: public), path (default: year/month/day)
2. Client can retrieve a file through public link
3. Client can retrieve a file by requesting access
- duration (in seconds)
4. Client can delete file permanently
5. Client should authenticate to upload a file
- api_key
- secret_key

## Technical Stack
1. Transport layer
- rest
- grpc
2. Database
- mysql
- mongo
3. Config
- system environment
