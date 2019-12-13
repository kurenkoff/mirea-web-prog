package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"mirea.com/web-prog/model"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/get", handlerGet).Methods(http.MethodGet)
	router.HandleFunc("/delete/{ids}", handleDelete).Methods(http.MethodDelete, http.MethodPost)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Access-Control-Allow-Origin"})
	corsObj := handlers.AllowedOrigins([]string{"*"})


	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headersOk, corsObj)(router)))
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	var (
		raw []byte
		err error
	)

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprintf(w, "error: %v", err.Error())
		} else {
			_, _ = fmt.Fprintf(w, "%v", string(raw))
		}
	}()

	raw, err = model.GetCollection()
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprintf(w, "error: %v", err.Error())
		}
	}()
	rawIDs := mux.Vars(r)["ids"]

	fmt.Println(rawIDs)
	model.Remove(strings.Split(rawIDs, ","))
}