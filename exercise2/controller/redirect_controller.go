package controller

import (
	"net/http"

	"github.com/krishnakumarkp/gophercises/exercise2/domain"
)

type RedirectController struct {
	Store domain.RedirectStore
}

func (rc RedirectController) PathHandler(w http.ResponseWriter, r *http.Request) (domain.Redirect, error) {
	path := r.URL.Path
	redirect, err := rc.Store.GetRedirectByPath(path)
	return redirect, err
}
