package http

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Listen starts the http server on the given port.
func Listen() {
	r := mux.NewRouter()

	s := r.PathPrefix("/index")

	s.HandlerFunc(createIndex).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}

func createIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey what up"))
}
