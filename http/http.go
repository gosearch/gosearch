package http

import (
	"net/http"
	"strconv"

	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gosearch/gosearch/service"
	"io"
	"io/ioutil"
)

const indexPath = "/index"

// Server holds the configuration for the HTTP server.
type Server struct {
	Index service.IndexService
}

// Listen starts the http server on the given port.
func (server *Server) Listen(port int) {
	router := mux.NewRouter()
	router.HandleFunc("/{index}/{id}", createIndex(server.Index)).Methods(http.MethodPost)
	router.HandleFunc("/{index}/{id}", getIndex(server.Index)).Methods(http.MethodGet)
	fmt.Println(http.ListenAndServe(":"+strconv.Itoa(port), router))
}

func createIndex(indexService service.IndexService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		index := vars["index"]
		id := vars["id"]
		data, err := bodyToJSON(r)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		_, err = indexService.Create(index, id, data)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func getIndex(indexService service.IndexService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		index := vars["index"]
		id := vars["id"]
		data, err := indexService.Get(index, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Println(data.GoString())
	}
}

func bodyToJSON(r *http.Request) (interface{}, error) {
	var data interface{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return nil, err
	}
	if err := r.Body.Close(); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func respondWithJSON(w http.ResponseWriter, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
