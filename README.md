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
    * [Separate test database](#test-database)
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
 - _(optional)_ [psql](https://www.postgresql.org/docs/13/app-psql.html)
 - _(optional)_ [redis-cli](https://redis.io/topics/rediscli)

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

todo

### Dependency injection

todo

### Test dependencies

todo

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
