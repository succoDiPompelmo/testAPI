package main

import (
	"github.com/gorilla/mux"
	"net/http"
    "encoding/json"
	"log"
	"fmt"
)

type App struct {
	Router *mux.Router
}

func (a *App) Init() {
	a.Router = mux.NewRouter()
    a.Router.HandleFunc("/product", GetProduct).Methods("GET")
}

func (a *App) Run(addr string) { 
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

type product struct {
	id string
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    
    if id == "" {
        respondWithError(w, 500, "CIAO")
        return
    }

	fmt.Println(id)

    respondWithJSON(w, http.StatusOK, product{id: string(id)})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}