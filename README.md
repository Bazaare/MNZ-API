# MNZ API

## Description
MNZ API is designed to receive a request from a Client and forward the request to the
Middleware Static Backend XML Service.

## Install
MNZ API could be installed in a number of ways. If I was deploying to a production environment I would do the following:
1. Use a CI system (e.g. GitLab CI alongside Go Releaser) to test, build, release and upload a binary for the required OS.
2. Use an Infrastructure as Code tool (Puppet, SaltStack, Ansible) to download the binary from the artifact store
3. Have the same IAC tool deploy the binary into a Docker container, exposing any required ports

## Configuration
MNZ API only has one command line flag:

`--insecure`

This flag disables JWT authentication and will pass through all requests to the Middleware Static Backend XML Service.

## Example usage
To test out MNZ API there are a few options:

### Option 1:

Run the code directly and make requests via the CLI (with authentication) - This can be done by sitting in the root directory (which contains
the main.go file) and running:

`go run main.go`

From your command line you can then run the following to receive your authentication token:

`curl -X POST localhost:8080/xml/login`

To retrieve the data from the Middleware Static Backend XML Service, you use the following command:

`curl -X GET localhost:8080/xml/1 -H "Authorization: Bearer <Token Here>"`

### Option 2

Run the code directly and make requests via the browser (without authentication) - This can be done by sitting in the root directory (which contains
the main.go file) and running:

`go run main.go --insecure`

To retrieve the data from the Middleware Static Backend XML Service, navigate to the following URL in the browser:

`localhost:8080/xml/1`


### Option 3

Build your required binary to run the application and make requests via the CLI (with authentication).

Build your binary:

#### Linux
`env GOOS=linux GOARCH=amd64 go build`

#### MacOS
`env GOOS=darwin GOARCH=amd64 go build`

#### Windows (Untested)
`env GOOS=windows GOARCH=amd64 go build`

Run the binary:

`./MNZ`

Request your authentication token:

`curl -X POST localhost:8080/xml/login`

Request your data from the Middleware Static Backend XML Service:

`curl -X GET localhost:8080/xml/1 -H "Authorization: Bearer <Token Here>"`




### Option 3

Build your required binary to run the application and make requests via the browser (without authentication)

Build your binary:

#### Linux
`env GOOS=linux GOARCH=amd64 go build`

#### MacOS
`env GOOS=darwin GOARCH=amd64 go build`

#### Windows (Untested)
`env GOOS=windows GOARCH=amd64 go build`

Run the binary:

`./MNZ`

To retrieve the data from the Middleware Static Backend XML Service, navigate to the following URL in the browser:

`localhost:8080/xml/1`