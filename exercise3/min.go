package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishnakumarkp/gophercises/exercise3/controller"
	"github.com/krishnakumarkp/gophercises/exercise3/jsonstore"
)

func main() {

	fileName := "./data/gopher.json"

	storyStore, err := jsonstore.NewStore(fileName)

	if err != nil {
		log.Fatal("Couldnot open story file!")
	}

	storyController := controller.StoryController{
		Store: storyStore,
	}

	router := mux.NewRouter()

	router.Handle("/story", storyController.ResponseHandler())
	router.Handle("/story/{arc}", storyController.ResponseHandler())

	http.ListenAndServe(":80", router)

}
