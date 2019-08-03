# Pre-Requisites

Install the latest stable version of Node and Go version 1.16 or higher, and then install the below Go programs. _Ensure the GOBIN directory is added to your PATH_.

```text
$ go install github.com/golang/mock/mockgen
$ go install github.com/google/wire/cmd/wire
```

# Build

Build the user interface:

```text
$ pushd web
$ npm install
$ npm run build
$ popd
```

Build the server and command line tools:

```text
$ go generate ./...
$ go build -o release/{{toLower name}}
```

# Test

Execute the unit tests:

```text
$ go generate ./...
$ go test -v -cover ./...
```

# Run

This {{toLower project}} supports all operating systems and architectures supported by Go.  This means you can build and run the system on your machine; docker containers are not required for local development and testing.

Start the server at `localhost:3000`

```text
$ release/{{toLower name}} server
```

# User Interface

This {{toLower project}} includes a simple user interface for interacting with the system. When you run the application, you can access the user interface by navigating to `http://localhost:3000` in your browser.

# Command Line

This {{toLower project}} includes simple command line tools for interacting with the system. Please remember that you must start the server before you can execute commands.

Register a new user:

```text
$ release/{{toLower name}} register
```

Login to the application:

```text
$ release/{{toLower name}} login
```

Logout from the application:

```text
$ release/{{toLower name}} logout
```

View your account details:

```text
$ release/{{toLower name}} account
```

Generate a peronsal access token:

```text
$ release/{{toLower name}} token
```

Create a {{toLower project}}:

```text
$ release/{{toLower name}} {{toLower project}} create <name>
```

List {{toLower project}}s:

```text
$ release/{{toLower name}} {{toLower project}} ls
```

Debug and output http responses from the server:

```text
$ DEBUG=true release/{{toLower name}} {{toLower project}} ls
```

View all commands:

```text
$ release/{{toLower name}} --help
```
