package gogenericshandler

import (
	"net/http"
)

type requester[Request any] interface {
	Build(*http.Request) Request
}

type respondor[Response any] interface {
	Render(http.ResponseWriter, *http.Request, Response, error)
}

type Handler[Request, Response any] interface {
	requester[Request]
	respondor[Response]

	Handle(Request) (Response, error)
}

func NewHandle[Request, Response any](handler Handler[Request, Response]) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request := handler.Build(r)
		response, err := handler.Handle(request)
		handler.Render(w, r, response, err)
	}
}
