```
├── api
│   ├── example
│   │   └── example.proto
│   └── openapiv2
│       └── example
│           └── example.swagger.json
├── cmd
│   └── grpc_server
│       └── main.go
│   ├── api
│   │   ├── grpc
│   │   │   ├── converters
│   │   │   ├── example
│   │   │   └── interceptors
│   │   └── http
│   │       ├── converters
│   │       ├── datatransfers
│   │       │   ├── request
│   │       │   └── response
│   │       ├── handlers
│   │       │   └── example
│   │       ├── middlewares
│   │       └── routes
│   │           └── example
│   ├── app
│   ├── config
│   ├── constants
│   ├── model
│   ├── providers
│   │   ├── di_provider
│   │   └── web_auth_provider
│   │       ├── keycloak_redis
│   │       ├── model
│   │       └── provider.go
│   ├── repository
│   │   ├── example
│   │   └── repository.go
│   ├── service
│   │   ├── example
│   │   └── service.go
│   └── utils
│       └── helpers
├── pkg
│   ├── api
│   │   ├── dart
│   │   │   └── example
│   │   ├── golang
│   │   │   └── example
│   │   └── typescript
│   │       └── example
│   └── vendor
│       ├── googleapis
│       └── grpc-gateway
├── docker-compose.yml
├── Makefile
├── README.MD
├── go.mod
├── go.sum
├── grpc.http
├── http-client.env.json
├── internal
├── openapitools.json
└── realms
    └── example-realm.json
```