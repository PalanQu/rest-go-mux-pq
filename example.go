package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type myApp struct {
	Router *mux.Router
}

var tmp int

func (a *myApp) addRouter(w http.ResponseWriter, r *http.Request) {
	tmp++
	fmt.Fprintf(w, "tmp++")
}

func (a *myApp) listRouter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's %d now", tmp)
}

func main() {
	a := myApp{}
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/add", a.addRouter).Methods("GET")
	a.Router.HandleFunc("/list", a.listRouter).Methods("GET")
	a.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
	})

	http.ListenAndServe(":8080", a.Router)
}
