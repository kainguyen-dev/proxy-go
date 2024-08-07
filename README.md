# Proxy Service

`Proxy-service`, is a centralized service with primary concerns are managing access from the outside
world to internal microservices:

    - Permission
    - Rate limiting
    - Routing
    - Authentication
    - Manage API key

Database:

    - Redis
    - PostGres

### 1. Creating project:
```sh
go mod init svc/proxy-service
curl https://raw.githubusercontent.com/gin-gonic/examples/master/basic/main.go > main.go
# Install gin
go get -u github.com/gin-gonic/gin
# Run 
go run main.go
```

### 2. Project structure:

### 3. Technical note: