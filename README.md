JSONpipe
========
This is a fork of [ARolek/jsonpipe](https://github.com/ARolek/jsonpipe) which handles registered actions with a chain of handlers to allow for robust and modular handling of requests.  A simple example is in the example folder that has three handlers in a chain.  You can run the example by executing the following command from the project root:

`go run example/example.go`

then in a separate shell window:

`go run example/example_client.go`

Also restructured connection management to handle hundreds of concurrent connections safely.

TODO
========
 * Write some unit tests
 * More docs/examples
