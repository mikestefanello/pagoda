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
* [Sessions](#sessions)
* [Authentication](#authentication)
  * [Login / Logout](#login--logout)
  * [Forgot password](#forgot-password)
  * [Registration](#registration)
  * [Authenticated user](#authenticated-user)
    * [Middleware](#middleware)
* [Routes](#routes)
  * [Custom middleware](#custom-middleware)
  * [Controller / Dependencies](#controller--dependencies)
  * [Patterns](#patterns)
  * [Custom middleware](#custom-middleware)
  * [Testing](#testing)
    * [HTTP server](#http-server)
    * [Request / Request helpers](#request--response-helpers)
    * [Goquery](#goquery)
* [Controller](#controller)
  * [Page](#page)
  * [Flash messaging](#flash-messaging)
  * [Pager](#pager)
  * [CSRF](#csrf)
  * [Automatic template parsing](#automatic-template-parsing)
  * [Cached responses](#cached-responses)
    * [Cache tags](#cache-tags)
    * [Cache middleware](#cache-middleware)
  * [Data](#data)
  * [Forms](#forms)
    * [Submission processing](#submission-processing)
    * [Inline validation](#inline-validation)
  * [Headers](#headers)
  * [Status code](#status-code)
  * [Metatags](#metatags)
  * [HTMX support](#htmx-support)
  * [Rendering the page](#rendering-the-page)
* [Template renderer](#template-renderer)
  * [Caching](#)
  * [Hot-reload for development](#)
* [Funcmap](#)
* [Cache](#cache)
* [Static files](#static-files)
  * Cache control headers
  * Cache-buster
* [Email](#email)
* [HTTPS](#https)
* [Logging](#logging)
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

Sessions are provided and handled via [Gorilla sessions](https://github.com/gorilla/sessions) and configured as middleware in the router located at `routes/router.go`. Session data is currently stored in cookies but there are many [options](https://github.com/gorilla/sessions#store-implementations) available if you wish to use something else.

Here's a simple example of loading data from a session and saving new values:

```go
func SomeFunction(ctx echo.Context) error {
	sess, err := session.Get("some-session-key", ctx)
	if err != nil {
		return err
	}
	sess.Values["hello"] = "world"
	sess.Values["isSomething"] = true
	return sess.Save(ctx.Request(), ctx.Response())
}
```

## Authentication

Included are standard authentication features you expect in any web application. Authentication functionality is bundled as a _Service_ within `services/AuthClient` and added to the `Container`. If you wish to handle authentication in a different manner, you could swap this client out or modify it as needed.

Authentication currently requires [sessions](#sessions) and the session middleware.

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

### Registration

The actual registration of a user is not handled within the `AuthClient` but rather just by creating a `User` entity. When creating a user, use `HashPassword()` to create a hash of the user's password, which is what will be stored in the database.

A route is provided for the user to register at `user/register`.

### Authenticated user

The `AuthClient` has two methods available to get either the `User` entity or the ID of the user currently logged in for a given request. Those methods are `GetAuthenticatedUser()` and `GetAuthenticatedUserID()`.

#### Middleware

Registered for all routes is middleware that will load the currently logged in user entity and store it within the request context. The middleware is located at `middleware.LoadAuthenticatedUser()` and, if authenticated, the `User` entity is stored within the context using the key `context.AuthenticatedUserKey`.

If you wish to require either authentication or non-authentication for a given route, you can use either `middleware.RequireAuthentication()` or `middleware.RequireNoAuthentication()`.

## Routes

The router functionality is provided by [Echo](https://echo.labstack.com/guide/routing/) and constructed within via the `BuildRouter()` function inside `routes/router.go`. Since the _Echo_ instance is a _Service_ on the _Container_ which is passed in to `BuildRouter()`, middleware and routes can be added directly to it.

### Custom middleware

By default, a middleware stack is included in the router that makes sense for most web applications. Be sure to review what has been included and what else is available within _Echo_ and the other projects mentioned.

A `middleware` package is included which you can easily add to along with the custom middleware provided.

### Controller / Dependencies

The `Controller`, which is described in a section below, serves two purposes for routes:

1) It provides base functionality which can be embedded in each route, most importantly `Page` rendering (described in the `Controller` section below)
2) It stores a pointer to the `Container`, making all _Services_ available within your route

While using the `Controller` is not required for your routes, it will certainly make development easier.

See the following section for the proposed pattern.

### Patterns

These patterns are not required, but were designed to make development as easy as possible.

To declare a new route that will have methods to handle a GET and POST request, for example, start with a new _struct_ type, that embeds the `Controller`:

```go
type Home struct {
	controller.Controller
}

func (c *Home) Get(ctx echo.Context) error {}

func (c *Home) Post(ctx echo.Context) error {}
```

Then create the route and add to the router:

```go
	home := Home{Controller: controller.NewController(c)}
	g.GET("/", home.Get).Name = "home"
	g.POST("/", home.Post).Name = "home.post"
```

Your route will now have all methods available on the `Controller` as well as access to the `Container`. It's not required to name the route methods to match the HTTP method.

**It is highly recommended** that you name your routes. Most methods on the back and frontend leverage the route name and parameters in order to generate URLs.

### Testing

Since most of your web application logic will live in your routes, being able to easily test them is important. The following aims to help facilitate that.

The test setup and helpers reside in `routes/router_test.go`.

Only a brief example of route tests were provided in order to highlight what is available. Adding full tests did not seem logical since these routes will most likely be changed or removed in your project.

#### HTTP server

When the route tests initialize, a new `Container` is created which provides full access to all of the _Services_ that will be available during normal application execution. Also provided is a test HTTP server with the router added. This means your tests can make requests and expect responses exactly as the application would behave outside of tests. You do not need to mock the requests and responses.

#### Request / Response helpers

With the test HTTP server setup, test helpers for making HTTP requests and evaluating responses are made available to reduce the amount of code you need to write. See `httpRequest` and `httpResponse` within `routes/router_test.go`.

Here is an example how to easily make a request and evaluate the response:

```go
func TestAbout_Get(t *testing.T) {
  doc := request(t).
    setRoute("about").
    get().
    assertStatusCode(http.StatusOK).
    toDoc()
}
```

#### Goquery

A helpful, included package to test HTML markup from HTTP responses is [goquery](https://github.com/PuerkitoBio/goquery). This allows you to use jQuery-style selectors to parse and extract HTML values, attributes, and so on.

In the example above, `toDoc()` will return a `*goquery.Document` created from the HTML response of the test HTTP server.

Here is a simple example of how to use it, along with [testify](https://github.com/stretchr/testify) for making assertions:

```go
h1 := doc.Find("h1.title")
assert.Len(t, h1.Nodes, 1)
assert.Equal(t, "About", h1.Text())
```

## Controller

As previously mentioned, the `Controller` acts as a base for your routes, though it is optional. It stores the `Container` which houses all _Services_ (_dependencies_) but also a wide array of functionality aimed at allowing you to build complex responses with ease and consistency.

### Page

The `Page` is the major building block of your `Controller` responses. It is a _struct_ type located at `controller/page.go`. The concept of the `Page` is that it provides a consistent structure for building responses and transmitting data and functionality to the templates.

All example routes provided construct and _render_ a `Page`. It's recommended that you review both the `Page` and the example routes as they try to illustrate all included functionality.

As you develop your application, the `Page` can be easily extended to include whatever data or functions you want to provide to your templates.

Initializing a new page is simple:

```go
func (c *Home) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
}
```

Using the `echo.Context`, the `Page` will be initialized with the following fields populated:

- `Context`: The passed in _context_
- `ToURL`: A function the templates can use to generate a URL with a given route name and parameters
- `Path`: The requested URL path
- `URL`: The requested URL
- `StatusCode`: Defaults to 200
- `Pager`: Initialized `Pager` (see below)
- `RequestID`: The request ID, if the middleware is being used
- `IsHome`: If the request was for the homepage
- `IsAuth`: If the user is authenticated
- `AuthUser`: The logged in user entity, if one
- `CSRF`: The CSRF token, if the middleware is being used
- `HTMX.Request`: Data from the HTMX headers, if HTMX made the request (see below)

### Flash messaging

While flash messaging functionality is provided outside of the `Controller` and `Page`, within the `msg` package, it's really only used within this context.

Flash messaging requires that [sessions](#sessions) and the session middleware are in place since that is where the messages are stored.

#### Creating messages

There are four types of messages, and each can be created as follows:
- Success: `msg.Success(ctx echo.Context, message string)`
- Info: `msg.Info(ctx echo.Context, message string)`
- Warning: `msg.Warning(ctx echo.Context, message string)`
- Danger: `msg.Danger(ctx echo.Context, message string)`

The _message_ string can contain HTML.

#### Rendering messages

When a flash message is retrieved from storage in order to be rendered, it is deleted from storage so that it cannot be rendered again.

The `Page` has a method that can be used to fetch messages for a given type from within the template: `Page.GetMessages(typ msg.Type)`. This is used rather than the _funcmap_ because the `Page` contains the request context which is required in order to access the session data. Since the `Page` is the data destined for the templates, you can use: `{{.GetMessages "success"}}` for example.

To make things easier, a template _component_ is already provided, located at `templates/components/messages.gohtml`. This will render all messages of all types simply by using `{{template "messages" .}}` either within your page or layout template.

### Pager

A very basic mechanism is provided to handle and facilitate paging located in `controller/pager.go`. When a `Page` is initialized, so is a `Pager` at `Page.Pager`. If the requested URL contains a `page` query parameter with a numeric value, that will be set as the page number in the pager.

During initialization, the _items per page_ amount will be set to the default, controlled via constant, which has a value of 20. It can be overridden by changing `Pager.ItemsPerPage` but should be done before other values are set in order to not provide incorrect calculations.

Methods include:

- `SetItems(items int)`: Set the total amount of items in the entire result-set
- `IsBeginning()`: Determine if the pager is at the beginning of the pages
- `IsEnd()`: Determine if the pager is at the end of the pages
- `GetOffset()`: Get the offset which can be useful is constructing a paged database query

There is currently no template (yet) to easily render a pager.

### CSRF

By default, all non GET requests will require a CSRF token be provided as a form value. This is provided by middleware and can be adjusted or removed in the router.

The `Page` will contain the CSRF token for the given request. There is a CSRF helper component template which can be used to easily render a hidden form element in your form which will contain the CSRF token and the proper element name. Simply include `{{template "csrf" .}}` within your form.

### Automatic template parsing

Dealing with templates can be quite tedious and annoying so the `Page` aims to make it as simple as possible with the help of the [template renderer](#template-renderer). To start, templates for _pages_ are grouped in the following directories within the `templates` directory:

- `layouts`: Base templates that provide the entire HTML wrapper/layout. This template should include a call to `{{template "content" .}}` to render the content of the `Page`.
- `pages`: Templates that are specific for a given route/page. These must contain `{{define "content"}}{{end}}` which will be injected in to the _layout_ template.
- `components`: A shared library of common components that the layout and base template can leverage.

Specifying which templates to render for a given `Page` is as easy as:

```go
page.Name = "home"
page.Layout = "main"
```

That alone will result in the following templates being parsed and executed when the `Page` is rendered:

1) `layouts/main.gohtml` as the base template
2) `pages/home.gohtml` to provide the `content` template for the layout
3) All template files located within the `components` directory
4) The entire [funcmap](#funcmap)

The [template renderer](#template-renderer) also provides caching and local hot-reloading.

### Cached responses

A `Page` can have cached enabled just by setting `Page.Cache.Enabled` to `true`. The `Controller` will automatically handle caching the HTML output, headers and status code. Cached pages are stored using a key that matches the full request URL and [middleware](#cache-middleware) is used to serve it on matching requests.

By default, the cache expiration time will be set according to the configuration value located at `Config.Cache.Expiration.Page` but it can be set per-page at `Page.Cache.Expiration`.

#### Cache tags

You can optionally specify cache tags for the `Page` by setting a slice of strings on `Page.Cache.Tags`. This provides the ability to build in cache invalidation logic in your application driven by events such as entity operations, for example.

The cache client on the `Container` is currently handled by [gocache](https://github.com/eko/gocache) which makes it easy to perform operations such as tag-invalidation, for example:

```go
c.Cache.Invalidate(ctx, store.InvalidateOptions{
    Tags: []string{"my-tag"},
})
```

#### Cache middleware

Cached pages are served via the middleware `ServeCachedPage()` in the `middleware` package.

The cache is bypassed if the requests meet any of the following criteria:
1) Is not a GET request
2) Is made by an authenticated user

Cached pages are looked up for a key that matches the exact, full URL of the given request.

### Data

The `Data` field on the `Page` is of type `interface{}` and is what allows your route to pass whatever it requires to the templates, alongside the `Page` itself.

### Forms

The `Form` field on the `Page` is similar to the `Data` field in that it's an `interface{}` type but it's meant to store a struct that represents a form being rendered on the page.

An example of this pattern is:

```go
type ContactForm struct {
    Email      string `form:"email" validate:"required,email"`
    Message    string `form:"message" validate:"required"`
    Submission controller.FormSubmission
}
```

Then in your page:

```go
page := controller.NewPage(ctx)
page.Form = ContactForm{}
```

How the _form_ gets populated with values so that your template can render them is covered in the next section.

#### Submission processing

Form submission processing is made extremely simple by leveraging functionality provided by [Echo binding](https://echo.labstack.com/guide/binding/), [validator](https://github.com/go-playground/validator) and the `FormSubmission` struct located in `controller/form.go`.

Using the example form above, these are the steps you would take within the _POST_ callback for your route:

Start by storing a pointer to the form in the conetxt so that your _GET_ callback can access the form values, which will be showed at the end:
```go
var form ContactForm
ctx.Set(context.FormKey, &form)
```

Parse the input in the POST data to map to the struct so it becomes populated:
```go
if err := ctx.Bind(&form); err != nil {
    // Something went wrong...
}
```

Process the submissions which uses [validator](https://github.com/go-playground/validator) to check for validation errors:
```go
if err := form.Submission.Process(ctx, form); err != nil {
    // Something went wrong...
}
```

Check if the form submission has any validation errors:
```go
if !form.Submission.HasErrors() {
    // All good, now execute something!
}
```

In the event of a validation error, you most likely want to re-render the form with the values provided and any error messages. Since you stored a pointer to the _form_ in the context in the first step, you can first have the _POST_ handler call the _GET_:
```go
if form.Submission.HasErrors() {
	return c.Get(ctx)
}
```

Then, in your _GET_ handler, extract the form from the context so it can be passed to the templates:
```go
page := controller.NewPage(ctx)
page.Form = ContactForm{}

if form := ctx.Get(context.FormKey); form != nil {
    page.Form = form.(*ContactForm)
}
```

And finally, your template:
```
<input id="email" name="email" type="email" class="input" value="{{.Form.Email}}">
```

#### Inline validation

The `FormSubmission` makes inline validation easier because it will store all validation errors in a map, keyed by the form struct field name. It also contains helper methods that your templates can use to provide classes and extract the error messages.

While [validator](https://github.com/go-playground/validator) is an incredible package that is used to validate based on struct tags, the downside is that the messaging, by default, is not very human-readable or easy to override. Within `FormSubmission.setErrorMessages()` the validation errors are converted to more readable messages based on the tag that failed validation. Only a few tags are provided as an example, so be sure to expand on that as needed.

To provide the inline validation in your template, there are two things that need to be done.

First, include a status class on the element so it will highlight green or red based on the validation:
```
<input id="email" name="email" type="email" class="input {{.Form.Submission.GetFieldStatusClass "Email"}}" value="{{.Form.Email}}">
```

Second, render the error messages, if there are any for a given field:
```
{{template "field-errors" (.Form.Submission.GetFieldErrors "Email")}}
```

### Headers
### Status code
### Metatags
### HTMX support
### Rendering the page