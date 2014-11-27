package jsonpipe

type Response struct {
	RequestId string      `json:"reqId"`
	Error     error       `json:"error,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}
