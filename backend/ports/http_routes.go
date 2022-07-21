package ports

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *HttpServer) router() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/status", s.GetStatusHandler)

	return router
}
