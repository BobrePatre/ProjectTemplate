## 📚 ProjectTemplate: A Comprehensive Golang Microservice Template ✨

**Welcome to ProjectTemplate!** 🎉 This repository houses a robust and well-structured Golang microservice template designed to jumpstart your backend development journey. 🚀

**Table of Contents**

* 🗺️ Project Structure
* 🧰 Key Features
* 🛠️ Technologies Used
* 🤝 Contributing
* 📝 License
* 🔗 Links

---

## 🗺️ Project Structure

The project is neatly organized into various folders, each with its specific purpose:

* **api:**
    * **openapiv2:** Houses OpenAPI v2 specifications for HTTP API documentation.
    * **proto:** Contains Protocol Buffer definitions for gRPC services.
* **cmd:**
    * **server:** Entry point for the application.
* **internal:**
    * **app:** Initializes and runs core application components (logger, gRPC server, gRPC gateway, HTTP server, graceful shutdown).
    * **config:** Handles service configuration logic.
    * **constants:** Defines service-wide constants.
    * **delivery:** Encapsulates communication (API) components.
        * **http:** Implements HTTP (REST) communication logic.
            * **converters:** Functions for data conversion between DTOs and domain models.
            * **datatransfers:** Stores data transfer objects (DTOs).
            * **handlers:** Implements HTTP handler logic.
            * **middlewares:** Contains middleware for authentication, CORS, and logging.
            * **routes:** Configures routing (maps requests to handlers, defines endpoints, sets middleware, and registers routes).
        * **grpc:** Implements gRPC communication logic.
            * **converters:** Functions for converting Protocol Buffer models to domain models and vice versa.
            * **impl:** Contains gRPC server implementations.
            * **interceptors:** Implements interceptors for authentication, logging, and more.
    * **models:** Defines internal domain models for the service.
    * **providers:** Offers various application providers.
        * **di_provider:** Provides dependency injection initialization.
        * **web_auth_provider:** Provides authentication abstractions.
            * **keycloak_redis:** Implements web_auth_provider for Keycloak and Redis.
    * **repository:** Contains repositories for data access.
    * **service:** Implements business logic.
    * **utils:** Provides utilities and helper functions.
* **migrations:** Stores database migration scripts.
* **pkg:** Houses generated code (gRPC server for Golang, client libraries for Dart).
* **realms:** Contains realm exports (Keycloak authentication configuration).
* **buf.gen.yaml:** Configuration for generating client libraries.
* **buf.work.yaml:** Configuration for Protocol Buffer declaration locations.
* **config.json:** Application configuration.
* **http-client.env.json:** HTTP client configuration.
* **requests.http:** Contains requests for API testing.

---

## 🧰 Key Features

* **Clean Architecture:**  Adheres to clean architecture principles for maintainability and testability.
* **gRPC & HTTP Support:**  Offers both gRPC and HTTP APIs for flexible communication.
* **Dependency Injection:** Employs dependency injection for better code organization and testability.
* **Authentication & Authorization:** Provides authentication and authorization mechanisms using Keycloak and Redis.
* **Database Migrations:** Manages database schema changes with migrations.
* **Logging & Middleware:** Includes logging and middleware functionalities for better control and observability.
* **Code Generation:** Leverages code generation for gRPC clients and servers.
* **Testing:** Facilitates testing with pre-configured tools and example requests.

---

## 🛠️ Technologies Used

* **Golang:** The primary programming language.
* **gRPC:** For high-performance RPC communication.
* **Protocol Buffers:** For data serialization and service definition.
* **OpenAPI:** For HTTP API documentation.
* **Keycloak:** For authentication and authorization.
* **Redis:** For caching and session management.
* **PostgreSQL:** As the database solution.

--- 

## 🤝 Contributing

Contributions are welcome! Please refer to the contributing guidelines for more information.

## 📝 License

This project is licensed under the [MIT License](LICENSE).

## 🔗 Links

* **Repository:** https://github.com/BobrePatre/ProjectTemplate 
