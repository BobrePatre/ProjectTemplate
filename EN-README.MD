```
├── delivery
│   ├── example
│   │   └── example.proto
│   └── opendeliveryv2
│       └── example
│           └── example.swagger.json
├── cmd
│   └── grpc_server
│       └── main.go
├── internal
│   ├── delivery
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
│   ├── delivery
│   │   ├── dart
│   │   │   └── example
│   │   ├── golang
│   │   │   └── example
│   │   └── typescript
│   │       └── example
│   └── vendor
│       ├── googledeliverys
│       └── grpc-gateway
├── docker-compose.yml
├── Makefile
├── README.MD
├── go.mod
├── go.sum
├── grpc.http
├── http-client.env.json
├── internal
├── opendeliverytools.json
└── realms
    └── example-realm.json
```


This project's structure is tailored for development using the Go programming language, offering support for both gRPC and REST delivery frameworks. Herein is a detailed explanation of each component within the project's architecture:

### Primary Directories

- **`delivery`**: Houses the delivery definitions.
    - **`example`**: Contains `example.proto`, the protobuf files for defining gRPC services.
    - **`opendeliveryv2/example`**: Includes `example.swagger.json`, detailing the REST delivery in Swagger format, which simplifies the generation of documentation and client libraries.

- **`cmd`**: Directory designated for executable files or entry points.
    - **`grpc_server`**: Comprises `main.go`, serving as the entry point to launch the gRPC server.

### Internal Structure (`internal`)

- **`delivery`**: Encapsulates the delivery handling logic.
    - **`grpc`**: A sub-directory for gRPC handlers
        - **`converters`**: Functions to convert between data formats.
        - **`example`**: Specific implementations for examples.
        - **`interceptors`**: gRPC interceptors that allow pre-processing of calls.
    - **`http`**: A sub-directory for HTTP handlers:
        - **`converters`**: Functions to convert between data formats.
        - **`datatransfers`**: Defines the data transfer objects (DTOs), including requests (`request`) and responses (`response`).
        - **`handlers`**: Handlers for HTTP requests.
        - **`middlewares`**: Middleware for processing HTTP requests.
        - **`routes`**: Defines the routes for the HTTP delivery.

- **`app`**: Core components (like gRPC server, HTTP server) and sets them up by integrating and utilizing dependencies from the DI provider.

- **`config`**: Stores the application's configurations.

- **`constants`**: Defines constants used across the application.

- **`model`**: Contains the data models.

- **`providers`**: Layer for supplying various services and components.
    - **`di_provider`**: Provider for dependency injection.
    - **`web_auth_provider`**: Providers for authentication and authorization.
        - **`keycloak_redis`**: An auth provider implementation adapted for use with Keycloak and Redis.
        - **`model`**: Data models utilized for authentication.
        - **`provider.go`**: An interface defining the auth provider's methods.

- **`repository`**: Contains abstractions for data storage operations.
    - **`example`**: An example repository.
    - **`repository.go`**: Common interfaces for repositories.

- **`service`**: Defines the services containing the business logic.
    - **`example`**: An example service.
    - **`service.go`**: Common interfaces for services.

- **`utils`**: Auxiliary utilities.
    - **`helpers`**: Helper functions.

### Additional Components and Configurations

- **`pkg`**: External libraries and generated client libraries for the project's delivery
    - **`delivery`**: Subdirectories for Dart, Go, TypeScript with client libraries and a gRPC server
    - **`vendor`**: External dependencies and libraries, including `googledeliverys` and `grpc-gateway`.

- **`docker-compose.yml`**, **`Makefile`**, **`go.mod`**, **`go.sum`**, **`grpc.http`**, **`http-client.env.json`**, **`opendeliverytools.json`**: Tools and configuration files for managing dependencies, building, deploying, and testing the project.

- **`realms`**: Configuration for realms (a feature of SSO Keycloak)
    - **`example-realm.json`**: An example of a realm configuration for KeyCloak.
