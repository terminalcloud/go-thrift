go-thrift
===

> A quick intro to thrift in golang

*Uses golang 1.4.2, thrift 0.9.2*
*The thrift-generated go has changed significantly between versions - be sure to use 0.9.2*

This repo is meant as a gentle introduction to thrift in golang. You should be familiar with both thrift and go beforehand:
- thrift.apache.org/tutorial/go

This intro is split into 5 commits, each of which introduces a new concept.
Any commit with a working server can be tested with a python client:

    thrift --gen py example.thrift
    ./client.py

### Commits

#### thriftfile

A simple thriftfile. One service extends another, and an exception is present.
The goal is to present as many esoteric cases in as little space as possible.

#### generate go thrift service

Notice the `//go:generate` statement at the top of `main.go`.
This is run using `go generate`, and will regenerate the go thrift files for our defined service.

Generated files are committed so that they can be consumed by others using `go get`.
It is the responsibility of the author to generate these files.

#### example/service handler

The handler is implemented, and plugged into a simple server.

#### multiple services

Another service is added to the thriftfile.

Each service lives on its own port. Services should be distinct enough that they should not share a port, barring other concerns. If you find two services needing to share a port, they should probably not be two services.

Server initialization is factored into its own file, so that each server can be restarted separately.

#### racy data access

The `count` function may be racy - it is refactored to use `chan`s.

If a service requires internal data access, it may become racy.
In that case, use `chan`s to avoid data corruption, adding a goroutine to the constructor to watch for requests.

In real scenarios, files with large implementations should also be split to maintain readability.

### NOTES

- Compiled thrift files go into a folder named after the thriftfile: example.thrift -> example/
- Service handlers live within a that services folder: example.Service -> example/service/handler.go

---

For a preconfigured environment to run through this tutorial, visit terminal.com.
