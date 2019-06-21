# Go Starter

This is a boilerplate Go project. It contains what most of your typical REST program
would (and more!):

* New Relic for instrumentation/monitoring
* Configuration, be it via .properties or 12 Factor
* Easy to use HTTP servers and clients
  * With a reference service already implemented and tested
* Structured logging

## What do I need installed

Only the Go related tooling. Get it at [the Go website](https://golang.org).

It doesn't hurt to get familiar with Go by following [the Go Tour](https://tour.golang.org) and [How to Write Go Code](https://golang.org/doc/code.html) and [Effective Go](https://golang.org/doc/effective_go.html).

## How to Use

1) Clone this repo
```
$ git clone github.com/nubunto/go-boilerplate
```

2) Change your package name to whatever you want
```
$ mv go-boilerplate myapp
```

4) Change the package name from the `user.go` file to your chosen package name

5) Change your module name in go.mod

6) Run:
```
$ cd go-boilerplate
$ go build -tags local ./cmd/http
```

The build tag `local` is a simple way to separate between environments. Take a look at `start.go` and `start_local.go`
to see how to implement your own build tags.

## Directory Structure

This is just a collection of libraries packaged in a well known directory structure
for Go projects.

This is how it is structured:

```
├── cmd
│   └── http
│       ├── aws.go
│       ├── config.go
│       ├── endpoints.go
│       ├── logger.go
│       ├── main.go
│       ├── responses.go
│       ├── routes.go
│       ├── services
│       │   ├── publisher.go
│       │   └── users.go
│       ├── start.go
│       └── start_local.go
├── go.mod
├── go.sum
├── README.md
└── user.go

```

The `cmd` directory contains your applications, each separated by a directory. `http` is a application
which has HTTP semantics, such as routing and instrumentation.

The root directory, usually named after your service, contains the code that is shared
between every application inside your `cmd` directory. This is where you would group your models,
which is the case of the `user.go` file, that contains, as you would expect, a type declaration
for a user of your system.

Feel free to split this into packages as much as you like when coding your service, but be aware that
you don't need to create a lot of packages in Go. Instead, try to group things by functionality and location
rather than thinking in packages.

## Files

The file names are pretty descriptive in what they contain, but here is a rundown:

- `aws.go`
  - Contains code related to AWS and it's services.
- `config.go`
  - Contains code that handles configuration: defaults, files, environment variables
- `endpoints.go`
  - Contains the code of your endpoints. This is where you implement your HTTP handlers.
- `logger.go`
  - Contains logging configuration.
- `responses.go`
  - Contains your HTTP responses.
- `routes.go`
  - Contains your routes, essentially tying your services and their dependencies to your app's routes.
- `main.go`
  - Initial wiring of your app: logging, configuration, routes, services, AWS, etc.
- `start.go`
  - Contains the `startApp` function, responsible for starting your app with the correct configuration.
- `start_local.go`
  - Contains the `startApp` function that works locally, i.e. not wiring the NewRelic agent.

## Libraries

This boilerplate contains the following libraries:

- [Viper](https://github.com/spf13/viper), for configuration (files, environment variables, etc)
- [Chi](https://github.com/go-chi/chi), a fast and flexible HTTP router
- [The NewRelic Go Agent](https://github.com/newrelic/go-agent) for instrumentation
- [Log15](https://github.com/inconshreveable/log15) for logging
- [pq](https://github.com/lib/pq), a high performance PostgreSQL driver
- [The Go AWS SDK](https://docs.aws.amazon.com/sdk-for-go/api/) for AWS services: SQS, SNS, etc.

This project contains examples on how to use these libraries, but not all use cases are covered.
Refer to the documentations of these libraries in order to better understand how to configure or do things not covered
by this documentation.

Go usually has a lot of important aspects of application development already available in the standard library,
such as HTTP, JSON, text (and HTML) templates. Remember to [skim through the Go stdlib](https://golang.org/pkg/#stdlib), you just might find
something there and avoid adding a third party dependency.

However, feel free to add or remove dependencies to your application as you see fit.
