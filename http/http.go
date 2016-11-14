package http

import (
	"net/http"
	"strconv"

	"strings"

	"github.com/gorilla/mux"
	"github.com/gosearch/gosearch/service"
)

const indexPath = "/index"

// Server holds the configuration for the HTTP server.
type Server struct {
	Index service.IndexService
}

// Listen starts the http server on the given port.
func (server *Server) Listen(port int) {
	r := mux.NewRouter()
	s := r.PathPrefix(indexPath)
	s.HandlerFunc(createIndex(server.Index)).Methods(http.MethodPost)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:" + strconv.Itoa(port),
	}
	srv.ListenAndServe()
}

func createIndex(indexService service.IndexService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.Replace(r.URL.Path, indexPath, "", 1)
		splitPath := strings.Split(path, "/")

		if len(splitPath) < 2 || splitPath[1] == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No index was specified."))
			return
		}
		// The component after the first '/'. Ignore the rest.
		index := splitPath[1]
		_, err := indexService.Create(index)
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Created index: " + index))
	}
}
