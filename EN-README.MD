## 📚 ProjectTemplate: Golang Microservice Template ✨

**Welcome to ProjectTemplate!** 🎉 This repository provides a robust and well-structured Golang microservice template to help you kickstart your backend development. 🚀

**Table of Contents**

* 🗺️ Project Structure
* 🧰 Key Features
* 🛠️ Technologies Used
* 🤝 Contribution

---

## 🗺️ Project Structure

The project is neatly organized into folders, each with a specific purpose:

* **api:**
    * **openapiv2:** Stores OpenAPI v2 specifications for HTTP API documentation.
    * **proto:** Contains Protocol Buffer definitions for gRPC services.
* **cmd:**
    * **server:** Entry point of the application.
* **internal:**
    * **app:** Initializes and runs the core application components (logger, gRPC server, gRPC gateway, HTTP server, graceful shutdown).
    * **config:** Handles service configuration logic.
    * **constants:** Defines constants used throughout the service.
    * **delivery:** Contains components responsible for communication (API).
        * **http:** Implements logic for communication over HTTP (REST).
            * **converters:** Functions for converting data between DTOs and domain models.
            * **datatransfers:** Stores data transfer objects (DTOs).
            * **handlers:** Implements the logic of HTTP handlers.
            * **middlewares:** Contains middleware for authentication, logging, etc.
            * **routes:** Configures routing (mapping requests to handlers, defining endpoints, setting middleware, and registering routes).
        * **grpc:** Implements logic for communication over gRPC.
            * **converters:** Functions for converting Protocol Buffer models to domain models and vice versa.
            * **impl:** Contains gRPC server implementations.
            * **interceptors:** Implements interceptors for authentication, logging, and other purposes.
    * **models:** Defines internal domain models for the service.
    * **providers:** Offers various providers for the application.
        * **di_provider:** Provides dependency injection initialization.
        * **web_auth_provider:** Provides abstractions for authentication.
            * **keycloak_redis:** Implements web_auth_provider for Keycloak and Redis.
    * **repository:** Contains repositories for working with data.
    * **service:** Implements business logic.
    * **utils:** Provides utilities and helper functions.
* **migrations:** Stores SQL database migrations.
* **pkg:** Contains generated code (gRPC server for Golang, client libraries for Dart).
* **realms:** Contains realm export (configuration for Keycloak authentication).
* **buf.gen.yaml:** Configuration for generating client libraries.
* **buf.work.yaml:** Configuration for the location of Protocol Buffer declarations.
* **config.json:** Application configuration.
* **http-client.env.json:** HTTP client configuration.
* **requests.http:** Contains requests for API testing.

---

## 🧰 Key Features

* **Clean Architecture:** Follows clean architecture principles for maintainability and testability.
* **gRPC and HTTP Support:** Offers both gRPC and HTTP APIs for flexible communication.
* **Dependency Injection:** Utilizes dependency injection for better code organization and testability.
* **Authentication and Authorization:** Provides authentication and authorization mechanisms using Keycloak and Redis.
* **Database Migrations:** Manages database schema changes with migrations.
* **Code Generation:** Leverages code generation for gRPC clients and servers.

---

## 🛠️ Technologies Used

* **Golang:** The primary programming language.
* **gRPC:** For high-performance RPC communication.
* **Protocol Buffers:** For data serialization and service definition.
* **OpenAPI:** For HTTP API documentation.
* **Keycloak:** For authentication and authorization.
* **Redis:** For caching.
* **PostgreSQL:** As the database solution.

--- 

## 🤝 Contribution

Your contributions are welcome! Feel free to open issues or pull requests.

## 📝 License

This project is licensed under the [MIT License](LICENSE).

## 🔗 Links

* **Repository:** https://github.com/BobrePatre/ProjectTemplate 
