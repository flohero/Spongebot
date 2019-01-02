package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func writeJson(w http.ResponseWriter, obj interface{}) {
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		internalError(w, err)
	}
}

func internalError(w http.ResponseWriter, err error) {
	writeError(w, err, 500)
}

func badRequest(w http.ResponseWriter, err error) {
	writeError(w, err, 400)
}

func notFound(w http.ResponseWriter, err error) {
	writeError(w, err, 404)
}

func forbidden(w http.ResponseWriter, err error) {
	writeError(w, err, 403)
}

func writeError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	w.Write([]byte(err.Error()))
}

func corsAndContentTypeHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if r.Method == "OPTIONS" {
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(200)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getIdFromPath(w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return -1, err
	}
	return id, nil
}

func created(w http.ResponseWriter) {
	w.WriteHeader(201)
}
