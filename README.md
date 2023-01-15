# Adomate API

Adomate Monolithic API.

## Introduction

This monolithic API is meant to break apart the web application into a backend and frontend.  This backend will provide data to the frontend as well as other web components in the future, ie: Mobile App.

This API is designed to be run in Kubernetes with a few helper tools. Ideally, we would be using Vault to feed Redis and Database credentials, and external secrets for other information, however, for our purpose this is not necessary so we will be using normal secrets. These secrets can be replaced with External Secrets, Sealed Secrets, etc. in the future or if necessary.

The API follows standard RESTful API designs. The API documentation is accessible through root, and the methods are standard:

- GET
- POST (create)
- PUT (replace)
- PATCH (update)
- DELETE

## Installation

Coming Soon...

### Local

To run locally, you'll need to setup a MySQL instance.

1. Download the desired release from github
2. Create a database, user and password in MySQL for the API to utilize
3. Rename .env.example and edit filling in the values as appropriate
4. Run `go run main.go` to bootstrap the database and start the webserver


### FAQ

1. How do I start the API automatically on boot?

    - Use a systemd service file, or utilize Docker/Containerd/Kubernetes.
        - For systemd, see the `systemd` directory for an example service file or visit [this link](https://www.digitalocean.com/community/tutorials/how-to-use-systemctl-to-manage-systemd-services-and-units) for more information.

2. How do I update the API?

    - Download the latest release, stop the API, replace the binary, and start the API.
    - If using Docker/Containerd/Kubernetes, follow the standard procedure for updating a container. The container is ephemeral, so data will not be lost.

3. How do I update the API without downtime?

    - Use a load balancer that supports zero downtime deployments.  This is not a requirement, but is recommended.
    - Use Kubernetes and utilize the rolling update strategy.

4. Is there Swagger/OpenAPI documentation available?

Swagger Documentation is currently WIP and is subject to change.

Yes, we have Swagger (OpenAPI 2.0) documentation available. When you start the API, the documentation can be found by visiting the `/swagger/` directory.  For example, if you are running the API on localhost on port 3000, you can visit <http://localhost:3000/swagger/index.html> to view the documentation.

## Unit Testing
Unit Testing is under development using the standard go testing pkg.  To run the tests, simply run `go test` in the root directory. I highly suggest setting gin to release and using verbose output for local testing, this can be done by setting the environment variable `GIN_MODE=release` and `go test -v` or as a single command `GIN_MODE=release go test -v`.

## Email Templates

The following are the coded email templates and available variables:

Coming Soon... example below.
- Register
    - FirstName
    - LastName

### Email Template Format

We use Go's templating engine to generate emails. More information can be found at:

- [Go html/template](https://pkg.go.dev/html/template)
- [Building Web Applications](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/07.4.html)

Emails send in HTML format using the Mailgun API.

### Email Template Functions

Coming Soon...
