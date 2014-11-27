package jsonpipe

type Handler func(response *Response, request *Request)

func (h Handler) And(next Handler) Handler {
	return Handler(func(response *Response, request *Request) {
		h(response, request)

		if response.Error != nil {
			// don't continue on if the previous was an error
			return
		}

		if next != nil {
			next(response, request)
		}
	})
}

func (h Handler) Then(next Handler) Handler {
	return Handler(func(response *Response, request *Request) {
		h(response, request)
		if next != nil {
			next(response, request)
		}
	})
}

func (h Handler) Run(request *Request) Response {
	response := Response{RequestId: request.RequestId}
	h(&response, request)
	return response
}
