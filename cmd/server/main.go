package main

import (
	"net/http"

	"github.com/blck-snwmn/gogenericshandler"
)

var _ gogenericshandler.Handler[string, string] = (*HelloWorldHandler)(nil)

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) Handle(request string) (string, error) {
	return request + " world\n", nil
}

func (h *HelloWorldHandler) Build(*http.Request) string {
	return "hello"
}
func (h *HelloWorldHandler) Render(w http.ResponseWriter, response string, err error) {
	w.Write([]byte(response))
}

func main() {
	hh := &HelloWorldHandler{}
	http.HandleFunc("/", gogenericshandler.NewHandle[string, string](hh))
	http.ListenAndServe(":8080", nil)
}
