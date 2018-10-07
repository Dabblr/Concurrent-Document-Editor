# Concurrent Document Editor
## Project Overview
### Deliverables
Our goal is to create a server-client system that:
* serves simple text documents over HTTP
* accepts updates to a document over HTTP
* stores and maintains revision history for a document
* allows multiple editors to work on the same document concurrently
* uses [Operational Transformation](https://en.wikipedia.org/wiki/Operational_transformation) to incorporate changes from multiple users into a single, newest version of the document
    * resolving conflicts as necessary

### Acceptance Criteria
At least initially, we can use a REST client (e.g. [Postman](https://www.getpostman.com/), [Insomnia](https://insomnia.rest/)) to test our sytem. Once we have the central pieces in place, we can create a JavaScript client with a GUI to perform all the edition operations.

#### Single User Use Case
- [ ] retrieve the latest version of a document with an `HTTP GET` request
- [ ] send updates for a specific document with an `HTTP POST` request
- [ ] confirm that the changes took effect by repeating the first step

#### Multi-User Use Case
The multi-user case can be "simulated" from a single machine:
- [ ] retrieve the latest version of a document **R1** with an `HTTP GET` request
- [ ] send an update **U1** with an `HTTP POST` request, identifying **R1** as the working copy
    - the latest version of the document on the server should now be **R2**
- [ ] send another update **U2** with an `HTTP POST` request, *also identifying **R1** as the working copy*
- [ ] confirm that the server applied **U2** on top of **R2**, even though we sent the update while working on **R1**

### Extensions
Time allowing, we'd like to consider extending our system to:
* store revision history for multiple documents
    * (this feature might be easy to sneak in earlier)
* support a basic document editing environment in the browser
* allow real-time edition of documents so that users see each other's editions as they happen
    * this would likely require that we iterate on our architecture (e.g. use web sockets instead of simple HTTP)

## Getting Started
* get this repo on your machine: `git clone git@github.com:Dabblr/Concurrent-Document-Editor.git`