<h1 align="center">
   Go Server
</h1>

<p align="center"><img src="/assets/golang.png" alt="golang" width="25%"/></p>

<h3 align="center">
  Clean Architecture + Wire + Gin + GraphQL(gqlgen) + MongoDB
</h3>

## Abstract

This project implements a server in Go using Clean Architecture principles for maintainable and scalable code. It leverages Gin for the HTTP web framework, gqlgen for GraphQL server implementation, Wire for dependency injection, and MongoDB as the database. JSON Web Tokens (JWT) are used for secure authentication.

## Clean Architecture

![Clean Architecture](./assets/clean.jpg)

Clean Architecture is designed to separate concerns, making the system easy to maintain and evolve. For more information, check out Uncle Bob's [Clean Architecture blog post](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

## Technologies

- **[Gin](https://github.com/gin-gonic/gin):** A high-performance web framework that is efficient and well-suited for building REST APIs.
- **[gqlgen](https://github.com/99designs/gqlgen):** A Go library for building GraphQL servers without any fuss.
- **[MongoDB](https://github.com/mongodb/mongo-go-driver):** A NoSQL database that offers high performance, high availability, and easy scalability.
- **[Wire](https://github.com/google/wire):** Wire is an automatic dependency injection tool for Go, which simplifies the process of wiring application components together.
- **[JWT](https://github.com/golang-jwt/jwt):** JSON Web Tokens are an open, industry standard RFC 7519 method for representing claims securely between two parties.

## Project Structure

```
.
├── cmd
└── pkg
├── config
├── constant
├── di
├── domain
├── infrastructure
│   ├── database
│   │   └── mongo
│   ├── graph
│   │   └── model
│   ├── logging
│   └── server
├── repository
│   ├── repository_interface
│   └── user_repository
└── usecase
├── usecase_interfaces
└── user_usecase

```

## Prerequisites

Before running the project, ensure you have the following installed:

- Go (version specified in `go.mod`)
- MongoDB server running locally or accessible remotely

## Running the Project

To run the project, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/jepbura/go-server.git
   ```

2. Change directory to the project:

   ```bash
   cd go-server
   ```

3. Install dependencies:

   ```bash
   go mod download
   ```

4. Change directory to the main:

   ```bash
   cd cmd
   ```

5. Run the main application:

   ```bash
   go run main.go
   ```

## Todo

- [x] Add Gin
- [x] Add GraphQL
- [x] Add MongoDB
- [x] Add Wire
- [ ] Write Test Cases
- [ ] Add JWT
- [ ] Add Gorilla WebSocket
- [ ] Add Redis
- [ ] Add Docker
- [ ] Add Kubernetes

## Acknowledgements

This project draws inspiration and code from the following repositories:

- [go-gin-clean-arch](https://github.com/thnkrn/go-gin-clean-arch)
- [clean-architecture-go](https://github.com/vidu171/clean-architecture-go)
- [gqlgen-clean-example](https://github.com/nutstick/gqlgen-clean-example)

Feel free to star and fork these repositories to show your support for the authors' work.
