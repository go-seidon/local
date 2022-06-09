# Documentation

## Functionality
1. Upload File
2. Retrieve File
3. Delete File

## Authentication
1. Basic Auth (username:password)
- Example:
```json
{
  "username": "service-user",
  "password": "some-secret"
}
```

- Header:
```json
{
  "Authorization": "Basic c2VydmljZS11c2VyOnNvbWUtc2VjcmV0"
}
```

2. API Key (api_key:secret_key)
- Example:
```json
{
  "api_key": "service-user",
  "secret_key": "some-secret"
}
```

- Header:
```json
{
  "X-Api-Key": "service-user",
  "X-Secret-Key": "some-secret"
}
```

## REST API
1. Upload File

2. Retrieve File

3. Delete File


## GRPC API
1. Upload File

2. Retrieve File

3. Delete File


## Database Schema: MySQL
1. File
```json
{
  "id": {
    "type": "int,auto-increment",
  },
  "unique_id": {
    "type": "varchar,required"
  },
  "name": {

  },
  "size": {

  },
  "extension": {
    
  }
}
```

## Database Schema: MongoDB
```json
{
  "_id": {
    "$oid": "628f70efd856995383294372"
  },
  "unique_id": "1815c960-6000-4ddf-955d-e93bf33d4c6c",
  "name": "jambi-kota.png",
  "mimetype": "image/png",
  "extension": "png",
  "size": 267839,
  "path": "file/default/2022/05/26/jambi-kota.png",
  "created_at": "2022-05-26T12:22:07.115Z",
  "updated_at": "2022-05-26T12:22:07.115Z"
}
```
