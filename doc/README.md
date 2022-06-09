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
