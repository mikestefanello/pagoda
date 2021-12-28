## (NAME) - Rapid, easy full-stack web development starter kit in Go

----
## Table of Contents
* [Introduction](#introduction)
    * [Overview](#overview)
    * [Motivation](#motivation)
    * [Foundation](#foundation)
      * [Backend](#backend)
      * [Frontend](#frontend)
      * [Storage](#storage)
    * [Screenshots](#screenshots)
* [Getting started](#getting-started)
  * [Dependencies](#dependencies)
  * [Start the application](#start-the-application)
  * [Running tests](#running-tests)
  * [Clients](#clients)
* [Service container](#service-container)
  * [Dependency injection](#dependency-injection)
  * [Test dependencies](#test-dependencies)
* [Configuration](#configuration)
    * [Environment overrides](#environment-overrides)
    * [Environments](#environments)
* [Database](#database)
    * [Auto-migrations](#auto-migrations)
    * [Separate test database](#separate-test-database)
* [ORM](#orm)
  * [Entity types](#entity-types)
  * [New entity type](#new-entity-type)
* [Authentication](#authentication)
  * [Login/Logout](#login-logout)
  * [Forgot password](#forgot-password)
  * [Registration](#registration)
* [Routes](#routes)
* [Controller / Page](#controller)
  * [Page](#)
  * [Flash messaging](#)
  * [Pager](#)
  * [CSRF](#)
  * [Automatic template parsing](#)
  * [Template caching](#)
  * [Template hot-reload for development](#)
  * [Funcmap](#)
  * [Cached responses](#)
  * [Cache tags](#)
  * [Inline form validation](#)
  * [HTMX support](#)
  * [Testing](#controller-testing)
    * [HTTP server](#)
    * [Response/request helpers](#)
    * [Goquery](#)
* [Template renderer](#template-renderer)
* [Cache](#cache)
  * Responses
  * Tags
* [Static files](#static-files)
  * Cache control headers
  * Cache-buster
* [Email](#email)
* [HTTPS](#https)
* [Roadmap](#roadmap)
* [Credits](#credits)

## Introduction

### Overview

(NAME) is not a framework but rather a base starter-kit for rapid, easy full-stack web development in Go, aiming to provide much of the functionality you would expect from a complete web framework as well as establishing patterns, procedures and structure for your web application.
 
todo

### Motivation

It started with [this post](https://news.ycombinator.com/item?id=29311761) on _Hacker News_, asking the community what the _simplest stack to build web apps in 2021_ is. After leaving PHP for Go over a year ago, I didn't have an answer for what I would use if I were to start building a web app tomorrow. If I was still using PHP, _Laravel_ would most likely be the easy answer, but there's nothing quite like that available for Go, especially in terms of adoption and maturity. For good reasons, the community also seems mostly opposed to mega-frameworks.

todo

### Foundation

While many great projects were used to build this, all of which are listed in the _credits_ section, the following provide the foundation of the back and frontend. It's important to note that you are **not required to use any of these**. Swapping any of them out will be relatively easy.

#### Backend

- [Echo](https://echo.labstack.com/): High performance, extensible, minimalist Go web framework.
- [Ent](https://entgo.io/): Simple, yet powerful ORM for modeling and querying data.

#### Frontend

Go server-side rendered HTML combined with the projects below enable you to create slick, modern UIs without writing any JavaScript or CSS.

- [HTMX](https://htmx.org/): Access AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML, using attributes, so you can build modern user interfaces with the simplicity and power of hypertext.
- [Alpine.js](https://alpinejs.dev/): Rugged, minimal tool for composing behavior directly in your markup. Think of it like jQuery for the modern web. Plop in a script tag and get going.
- [Bulma](https://bulma.io/): Provides ready-to-use frontend components that you can easily combine to build responsive web interfaces. No JavaScript dependencies.

#### Storage

- [PostgreSQL](https://www.postgresql.org/): The world's most advanced open source relational database.
- [Redis](https://redis.io/): In-memory data structure store, used as a database, cache, and message broker.

### Screenshots

todo

## Getting started

### Dependencies

Ensure the following are installed on your system:

 - [Go](https://go.dev/)
 - [Docker](https://www.docker.com/)
 - [Docker Compose](https://docs.docker.com/compose/install/)
 - [psql](https://www.postgresql.org/docs/13/app-psql.html) _(optional)_
 - [redis-cli](https://redis.io/topics/rediscli) _(optional)_

### Start the application

After checking out the repository, from within the root, start the Docker containers for the database and cache by executing `make up`.

Once that completes, you can start the application by executing `make run`. By default, you should be able to access the application in your browser at `localhost:8000`.

If you ever want to quickly drop the Docker containers and restart them in order to wipe all data, execute `make reset`.

### Running tests

To run all tests in the application, execute `make test`. This ensures that the tests from each package are not run in parallel. This is required since many packages contain tests that connect to the test database which is dropped and recreated automatically for each package.

### Clients

The following _make_ commands are available to make it easy to connect to the database and cache.

- `make db`: Connects to the primary database
- `make db-test`: Connects to the test database
- `make cache`: Connects to the cache

## Service container

The container is located at `services/container.go` and is meant to house all of your application's services and/or dependencies. It is easily extensible and can be created and initialized in a single call. The services currently included in the container are:

- Configuration
- Cache
- Database
- ORM
- Web
- Validator
- Authentication
- Mail
- Template renderer

A new container can be created and initialized via `services.NewContainer()`. It can be later shutdown via `Shutdown()`.

### Dependency injection

The container exists to faciliate easy dependency-injection both for services within the container as well as areas of your application that require any of these dependencies. For example, the container is passed to and stored within the `Controller`
 so that the controller and the route using it have full, easy access to all services.

### Test dependencies

It is common that your tests will require access to dependencies, like the database, or any of the other services available within the container. Keeping all services in a container makes it especially easy to initialize everything within your tests. You can see an example pattern for doing this [here](#environments).

## Configuration

The `config` package provides a flexible, extensible way to store all configuration for the application. Configuration is added to the _Container_ as a _Service_, making it accessible across most of the application. 

Be sure to review and adjust all of the default configuration values provided.

### Environment overrides

Leveraging the functionality of [envdecode](https://github.com/joeshaw/envdecode), all configuration values can be overridden by environment variables. Here is an example of what a configuration value looks like, each of which is a field on a struct:

```go
Port         uint16        `env:"HTTP_PORT,default=8000"`
```

The value for this field will be set to `8000`, the default, unless the `HTTP_PORT` environment variable is set, in which case the value of the variable will be used. This allows you to easily override configuration values per-environment.

### Environments

The configuration value for the current _environment_ (`Config.App.Environment`) is an important one as it can influence some behavior significantly (will be explained in later sections).

A helper function (`config.SwitchEnvironment`) is available to make switching the environment easy, but this must be executed prior to loading the configuration. The common use-case for this is to switch the environment to `Test` before tests are executed:

```go
func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnvironment(config.EnvTest)

	// Start a new container
	c = services.NewContainer()
	defer func() {
		if err := c.Shutdown(); err != nil {
			c.Web.Logger.Fatal(err)
		}
	}()

	// Run tests
	exitVal := m.Run()
	os.Exit(exitVal)
}
```

## Database

The database currently used is [PostgreSQL](https://www.postgresql.org/) but you are free to use whatever you prefer. If you plan to continue using [Ent](https://entgo.io/), the incredible ORM, you can check their supported databases [here](https://entgo.io/docs/dialects). The database-driver and client is provided by [pgx](github.com/jackc/pgx/v4) and included in the `Container`.

Database configuration can be found and managed within the `config` package.

### Auto-migrations

[Ent](https://entgo.io/) provides automatic migrations which are executed on the database whenever the `Container` is created, which means they will run when the application starts.

### Separate test database

Since many tests can require a database, this application supports a separate database specifically for tests. Within the `config`, the test database name can be specified at `Config.Database.TestDatabase`.

When a `Container` is created, if the [environment](#environments) is set to `config.EnvTest`, the database client will connect to the test database instead, drop the database, recreate it, and run migrations so your tests start with a clean, ready-to-go database. Another benefit is that after the tests execute in a given package, you can connect to the test database to audit the data which can be useful for debugging.

## ORM

As previously mentioned, [Ent](https://entgo.io/) is the supplied ORM. It can swapped out, but I highly recommend it. I don't think there is anything comparable for Go, at the current time. If you're not familiar with Ent, take a look through their top-notch [documentation](https://entgo.io/docs/getting-started). 

An Ent client is included in the `Container` to provide easy access to the ORM throughout the application.

Ent relies on code-generation for the entities you create to provide robust, type-safe data operations. Everything within the `ent` package in this repository is generated code for the two entity types listed below with the exception of the schema declaration.

### Entity types

The two included entity types are:
- User
- PasswordToken

### New entity type

While you should refer to their [documentation](https://entgo.io/docs/getting-started) for detailed usage, it's helpful to understand how to create an entity type and generate code. To make this easier, the `Makefile` contains some helpers.

1. Ensure all Ent code is downloaded by executing `make ent-install`.
2. Create the new entity type by executing `make ent-new name=User` where `User` is the name of the entity type. This will generate a file like you can see in `ent/schema/user.go` though the `Fields()` and `Edges()` will be left empty.
3. Populate the `Fields()` and optionally the `Edges()` (which are the relationships to other entity types).
4. When done, generate all code by executing `make ent-gen`.

The generated code is extremely flexible and impressive. An example to highlight this is one used within this application:

```go
entity, err := ORM.PasswordToken.
		Query().
		Where(passwordtoken.HasUserWith(user.ID(userID))).
		Where(passwordtoken.CreatedAtGTE(expiration)).
		All(ctx.Request().Context())
```

This executes a database query to return all _password token_ entities that belong to a user with a given ID and have a _created at_ timestamp field that is greater than or equal to a given time.

## Sessions

Sessions are provided and handled via [Gorilla sessions](https://github.com/gorilla/sessions) and configured in the router located at `routes/router.go`. Session data is currently stored in cookies but there are many [options](https://github.com/gorilla/sessions#store-implementations) available if you wish to use something else.

## Authentication

Included are standard authentication features you expect in any web application. Authentication functionality is bundled as a _Service_ within `services/AuthClient` and added to the `Container`. If you wish to handle authentication in a different manner, you could swap this client out or modify it as needed.

### Login / Logout

The `AuthClient` has methods `Login()` and `Logout()` to log a user in or out. To track a user's authentication state, data is stored in the session including the user ID and authentication status.

Prior to logging a user in, the method `CheckPassword()` can be used to determine if a user's password matches the hash stored in the database and on the `User` entity.

Routes are provided for the user to login and logout at `user/login` and `user/logout`.

### Forgot password

Users can reset their password in a secure manner by issuing a new password token via the method `GeneratePasswordResetToken()`. This creates a new `PasswordToken` entity in the database belonging to the user. The actual token itself, however, is not stored in the database for security purposes. It is only returned via the method so it can be used to build the reset URL for the email. Rather, a hash of the token is stored, using `bcrypt` the same package used to hash user passwords. The reason for doing this is the same as passwords. You do not want to store a plain-text value in the database that can be used to access an account.

Tokens have a configurable expiration. By default, they expire within 1 hour. This can be controlled in the `config` package. The expiration of the token is not stored in the database, but rather is used only when tokens are loaded for potential usage. This allows you to change the expiration duration and affect existing tokens.

Since the actual tokens are not stored in the database, the reset URL must contain the user's ID. Using that, `GetValidPasswordToken()` will load all non-expired _password token_ entities belonging to the user, and use `bcrypt` to determine if the token in the URL matches any of the stored hashes.

Once a user claims a valid password token, all tokens for that user should be deleted using `DeletePasswordTokens()`.

Routes are provided to request a password reset email at `user/password` and to reset your password at `user/password/reset/token/:uid/:password_token`.

  * [Registration](#registration)