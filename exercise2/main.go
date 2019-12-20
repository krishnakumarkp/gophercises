package main

import (
	"log"
	"net/http"

	"github.com/krishnakumarkp/gophercises/exercise2/controller"
	"github.com/krishnakumarkp/gophercises/exercise2/store"
)

func main() {
	datastore, err := store.New("./data/path.json")

	if err != nil {
		log.Fatal("Couldnot create datastore")
	}

	redirectController := controller.RedirectController{
		Store: datastore,
	}

	mux := http.NewServeMux()
	mux.Handle("/", controller.ResponseHandler(redirectController.PathHandler))
	http.ListenAndServe(":80", mux)
}
