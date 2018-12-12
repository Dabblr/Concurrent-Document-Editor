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
* [Postman](https://www.getpostman.com/apps) (optionally, for testing)

### Building & Running
See [Go Workspaces](https://golang.org/doc/code.html#Workspaces) for info on the directory structure assumed by tools like `go get`.
```bash
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

With the server running locally (on port `8080` by default), you can send a `cURL` request thus:
```shell
$ curl http://localhost:8080/files/1      # causes error message in server logs, since file with id=1 does not exist yet
```

### Testing
#### Unit Tests
Each package has its own set of tests. They can be run using the `go test` command. For example, to run the Operational Transformation tests:
```
$ cd operationaltransformation/
$ go test                         # use "-v" (verbose) option for more info
PASS
ok  	github.com/dabblr/Concurrent-Document-Editor/operationaltransformation	0.004s
```



#### API Tests
A collection of API tests -- `server_tests.postman_collection.json` -- is included in the `server/` directory. Open the collection by clicking on the `Import` menu in Postman and selecting the collection file (additional info on Postman collections [here](https://learning.getpostman.com/docs/postman/collections/managing_collections)). The tests validate various error and success cases for each of our endpoints:
* **POST** `/users`
  * Create a user, which is required in order to create a file
* **POST** `/files`
  * Create a file with the given name as its name and given user as its owner
* **GET** `/files/{id}`
  * Get the specified file
* **POST** `/files/{id}`
  * Update the specified file with the given changes, which are defined with respect to a specific revision
  * This is where divergence is resolved!

Here is an example of a valid body for updating the contents of file through a **POST** request to `/files/{id}`:
```
{
	"user":"user1",
	"id":{{id}},
	"revision_number":1,
	"name": "file1.txt",
	"changes":[
		{
			"type":"insert",
			"position":0,
			"value":"x"
		},
		{
			"type":"insert",
			"position":1,
			"value":"y"
		}
	]
}
```

