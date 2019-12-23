package main

import (
	"log"
	"net/http"

	util "github.com/krishnakumarkp/gophercises/exercise2/apputil"
	"github.com/krishnakumarkp/gophercises/exercise2/controller"
	"github.com/krishnakumarkp/gophercises/exercise2/store"
	"github.com/spf13/viper"
)

func main() {
	datastore, err := store.New(util.AppConfig.File)

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

func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	util.AppConfig.File = viper.GetString("redirect.File")

}
