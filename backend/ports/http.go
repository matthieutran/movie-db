package ports

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/matthieutran/movie-db/backend/app"
)

const version = "0.0.1"

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application) *HttpServer {
	return &HttpServer{
		app: app,
	}
}

func (s *HttpServer) StartHttpServer(wg sync.WaitGroup) func(host string, port int) error {
	return func(host string, port int) error {
		httpServer := &http.Server{
			Addr:         fmt.Sprintf("%s:%d", host, port),
			Handler:      s.router(),
			IdleTimeout:  time.Minute,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 30 * time.Second,
		}

		fmt.Println("Starting HTTP Server on port", port)

		err := httpServer.ListenAndServe()
		if err != nil {
			wg.Done()
		}
		return err
	}
}

type StatusResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func (s *HttpServer) GetStatusHandler(w http.ResponseWriter, r *http.Request) {
	currStatus := StatusResponse{
		Status:  "Available",
		Version: version,
	}

	js, err := json.MarshalIndent(currStatus, "", "\t")
	if err != nil {
		fmt.Println("Could not marshal status:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
