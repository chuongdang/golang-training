# User Service

## Requirements
- golang >= 1.11

## 1. How to run

```go
go run basic
```

## 2. How to test

You can either run these following commands or import them using https://www.getpostman.com/ to test

### 2.1 Get user list
```bash
curl -H 'X-Nordic-Token: user' 'http://127.0.0.1:6969/user'
```

### 2.1 Add user
```bash
curl -X POST 'http://127.0.0.1:6969/user?limit=10&page=1' -H 'X-Nordic-Token: admin' -d '{"role_id": 1, "name": "Test", "email": "test@gmail.com", "password": "123123"}'
```