package server

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/vlomel/vlomel_skillbox_diploma/internal/result_data"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res []byte
	res, _ = json.Marshal(result.GetResultData())
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(res); err != nil {
		log.Err(err).Msg("Ошибка записи ответа")

	}
}

func (s *Server) RunServer(port string) error {
	r := mux.NewRouter()
	r.HandleFunc("/", handleConnection).Methods("GET")

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
