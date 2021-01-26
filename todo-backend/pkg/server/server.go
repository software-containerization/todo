package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"todo-api/pkg/storage"

	"github.com/google/wire"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Server exposes the API
type Server struct {
	router *mux.Router
	repo   *storage.ItemRepository
}

// Set is the autoinjection set for Server
var Set = wire.NewSet(
	mux.NewRouter,
	NewServer,
	storage.Set,
)

// NewServer returns a new Server Object
func NewServer(r *mux.Router, re *storage.ItemRepository) *Server {
	return &Server{
		router: r,
		repo:   re,
	}
}

// Run starts the server on port
func (s *Server) Run(port string) {
	s.routes()
	log.Infof("Starting server on %s", port)

	srv := &http.Server{
		Handler:      s.router,
		Addr:         fmt.Sprintf(":%s", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func (s *Server) routes() {
	s.router.Use(AllowCors)
	s.router.HandleFunc("/api/health", HandleHealth).Methods("GET")
	s.router.HandleFunc("/api/item", s.handlePostItem()).Methods("POST", "OPTIONS")
	s.router.HandleFunc("/api/items", s.handleGetItems()).Methods("GET")
	s.router.HandleFunc("/api/item/{id:[0-9]+}", s.handleDeleteItem()).Methods("DELETE", "OPTIONS")
	s.router.HandleFunc("/api/item/{id:[0-9]+}", s.handleUpdateItem()).Methods("PUT", "OPTIONS")
}

func AllowCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Got %s request on %s", r.Method, r.RequestURI)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}

// HandleHealth handles the healthcheck HTTP request
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	if err != nil {
		log.Error(err)
	}
}
