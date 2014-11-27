package jsonpipe

type Request struct {
	Action    string                 `json:"action"`
	Error     error                  `json:"error,omitempty"`
	RequestId string                 `json:"reqId"`
	Data      map[string]interface{} `json:"data"`
}
