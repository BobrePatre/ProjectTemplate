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

Данная структура проекта предназначена для организации разработки на языке Go,
поддерживающей как gRPC, так и REST delivery. Вот подробное описание каждой из компонентов структуры проекта:

### Основные Каталоги

- **`delivery`**: Содержит определения delivery.
    - **`example`**: Включает `example.proto`, файлы protobuf для определения gRPC сервисов.
    - **`opendeliveryv2/example`**: Содержит `example.swagger.json`, описания REST delivery в формате Swagger, что облегчает
      генерацию документации и клиентских библиотек.

- **`cmd`**: Каталог для исполняемых файлов или точек входа.
    - **`grpc_server`**: Включает `main.go`, точку входа для запуска gRPC сервера.

### Внутренняя Структура (`internal`)

- **`delivery`**: Содержит логику обработки deliveryы
- 
- .
    - **`grpc`** Подкаталог для gRPC обработчиков
        - **`converters`**: Преобразователи между форматами данных.
        - **`example`**: Специфичные реализации для примеров.
        - **`interceptors`**: Интерсепторы для gRPC, позволяющие предварительно обрабатывать вызовы.
    - **`http`**: Подкаталог для HTTP обработчиков:
        - **`converters`**: Преобразователи между форматами данных.
        - **`datatransfers`**: Определения для объектов передачи данных (DTOs), включая запросы (`request`) и
          ответы (`response`).
        - **`handlers`**: Обработчики HTTP запросов.
        - **`middlewares`**: Промежуточное программное обеспечение для HTTP запросов.
        - **`routes`**: Определения маршрутов для HTTP delivery.

- **`app`**: Основные компоненты (gRPC server, HTTP server) и их конфигурация с вредрением зависимостей из di провайдера

- **`config`**: Хранит конфигурации приложения.

- **`constants`**: Определяет константы, используемые в приложении.

- **`model`**: Содержит модели данных.

- **`providers`**: Слой для предоставления различных сервисов и компонентов.
    - **`di_provider`**: Провайдер для внедрения зависимостей.
    - **`web_auth_provider`**: Провайдеры аутентификации и авторизации.
        - **`keycloak_redis`**: Реализация auth провайдера адаптированного для использования keycloak и redis
        - **`model`**: Модели данных, используемые для аутентификации.
        - **`provider.go`**: Интерфейс определяющий методы auth провайдера

- **`repository`**: Содержит абстракции для работы с хранилищем данных.
    - **`example`**: Пример репозитория.
    - **`repository.go`**: Общие интерфейсы для репозиториев.

- **`service`**: Определяет сервисы, содержащие бизнес-логику.
    - **`example`**: Пример сервиса.
    - **`service.go`**: Общие интерфейсы для сервисов.

- **`utils`**: Вспомогательные утилиты.
    - **`helpers`**: Вспомогательные функции.

### Дополнительные Компоненты и Конфигурации

- **`pkg`**: Внешние библиотеки и сгенерированные клиентские библиотеки для delivery проекта
    - **`delivery`**: Подкаталоги для Dart, Go, TypeScript с клиентскими библиотеками и grpc сервером
    - **`vendor`**: Внешние зависимости и библиотеки, включая `googledeliverys` и `grpc-gateway`.


- **`docker-compose.yml`**, **`Makefile`**, **`go.mod`**, **`go.sum`**, **`grpc.http`**, **`http-client.env.json`**, *
  *`opendeliverytools.json`**: Инструменты и файлы конфигурации для управления зависимостями, сборки, развертывания и
  тестирования проекта.

- **`realms`**: Конфигурации realm (Особенность SSO Keycloak)
    - **`example-realm.json`**: Пример конфигурации realm для KeyCloak