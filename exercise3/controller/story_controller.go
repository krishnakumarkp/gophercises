package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishnakumarkp/gophercises/exercise3/domain"
)

type StoryController struct {
	Store domain.StoryStore
}

func (sc StoryController) ResponseHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var arc string
		params := mux.Vars(r)

		arc = "intro"

		if val, ok := params["arc"]; ok {
			arc = val
		}

		story, err := sc.Store.GetByArc(arc)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "Storynot found!")
			return
		}

		renderTemplate(w, story)
	})
}

func renderTemplate(w http.ResponseWriter, data interface{}) {
	t, err := template.ParseFiles("template/story.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
