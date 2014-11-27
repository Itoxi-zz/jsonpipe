package main

import (
	"log"
	. "github.com/itoxi/jsonpipe"
)

func main() {

	handler := NewDoMagicHandler().And(NewLogHandler()).Then(NewEndItAllHandler())
	server := NewServer()
	server.Handle("foo", handler)
	server.ListenAndServe("0.0.0.0:8080")
}

func NewDoMagicHandler() Handler {
	return func(response *Response, request *Request) {
		//do some magic here and set a value on the response
		response.Data = "MAGIC INDEED HAPPENED"
	}
}

func NewLogHandler() Handler {
	return func(response *Response, request *Request) {
		log.Printf("Magic? %v", response.Data)
	}
}

func NewEndItAllHandler() Handler {
	return func(response *Response, request *Request) {
		log.Println("...The End")
	}
}
