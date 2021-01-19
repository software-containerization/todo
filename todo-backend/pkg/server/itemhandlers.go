package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-api/pkg/data"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func (s *Server) handleGetItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := s.repo.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(items)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (s *Server) handlePostItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if t := r.Header.Get("Content-Type"); t != "application/json;charset=utf-8" {
			http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
			return
		}

		var i data.Item

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&i)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = s.repo.Create(i)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) handleDeleteItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		err := s.repo.Delete(id)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) handleUpdateItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if t := r.Header.Get("Content-Type"); t != "application/json;charset=utf-8" {
			http.Error(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
			return
		}

		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 16)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var i data.Item

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err = decoder.Decode(&i)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		i.ID = id

		err = s.repo.Update(i)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
