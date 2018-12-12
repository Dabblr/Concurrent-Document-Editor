## Concurrent Document Editor
### Project Overview
This repo contains a self-contained prototype of a server-client shared document editor system that:
* serves simple text documents over HTTP
* allows a user to create a new text document with an HTTP request
* accepts updates to a document over HTTP
* stores and maintains revision history for a document
* allows multiple editors to work on the same document concurrently
* uses [Operational Transformation](https://en.wikipedia.org/wiki/Operational_transformation) to resolve and incorporate conflicting revisions of a document

### Prerequisites
* [SQLite](https://www.sqlite.org/download.html), which we use to track documents, changes, revisions, and users.
* [Golang](https://golang.org/doc/install), which we use to implement all components of the system

### Building & Running
See info on [Go Workspaces](https://golang.org/doc/code.html#Workspaces) for info on the directory structure assumed by tools like `go get`.
```
$ cd $GOPATH                              # $GOPATH env var has path to Go workspace
$ pwd
/Users/juan/code/go

$ mkdir -p src/github.com/dabblr/         # create new folder, into which we'll clone the repo
$ cd src/github.com/dabblr/
$ git clone https://github.com/Dabblr/Concurrent-Document-Editor.git

$ cd Concurrent-Document-Editor/server
$ go get                                  # install packages
$ go build                                # compile!

$ ./server
Starting real database
```

With the server running locally (on port `8080` by default), you can send a `cURL` thus:
```
$ curl http://localhost:8080/files/1      # causes error message in server logs, since file with id=1 does not exist yet
```
