package main

import (
	"net/http"

	"github.com/blck-snwmn/gogenericshandler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var _ gogenericshandler.Handler[string, string] = (*HogeHandler)(nil)

type HogeHandler struct{}

func (hh *HogeHandler) Handle(request string) (string, error) {
	return request + " world\n", nil
}

func (hh *HogeHandler) Build(*http.Request) string {
	return "hello"
}
func (hh *HogeHandler) Render(w http.ResponseWriter, response string, err error) {
	w.Write([]byte(response))
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", gogenericshandler.NewHandle[string, string](&HogeHandler{}))

	http.ListenAndServe(":8080", r)
}
