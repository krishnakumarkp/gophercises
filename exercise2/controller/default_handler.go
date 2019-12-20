package controller

import (
	"fmt"
	"net/http"

	"github.com/krishnakumarkp/gophercises/exercise2/domain"
)

func ResponseHandler(h func(http.ResponseWriter, *http.Request) (domain.Redirect, error)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		redirect, err := h(w, r)
		if err != nil {
			fmt.Fprintf(w, "Hello , you have hit %s \n", r.URL.Path)
		} else {
			http.Redirect(w, r, redirect.Url, 307)
		}
	})
}
