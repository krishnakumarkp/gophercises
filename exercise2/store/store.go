package store

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/krishnakumarkp/gophercises/exercise2/domain"
)

func New(filename string) (domain.RedirectStore, error) {
	var store domain.RedirectStore
	var err error
	extension := strings.ToLower(filepath.Ext(filename))
	fmt.Println(extension)
	switch extension {
	case ".json":
		store, err = NewJsonStore(filename)
	case ".yml":
		store, err = NewYamlStore(filename)
	case ".bolt":
		//store = store.NewBoltStore(filename)
	default:
		store = NewMapStore()
	}
	return store, err
}
