package store

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/krishnakumarkp/gophercises/exercise2/domain"
)

type JsonStore struct {
	store []domain.Redirect
}

func NewJsonStore(file string) (*JsonStore, error) {
	var jstore JsonStore
	jsonFile, err := os.Open(file)
	if err != nil {
		return &jstore, err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return &jstore, err
	}
	err = json.Unmarshal(byteValue, &jstore.store)
	if err != nil {
		return &jstore, err
	}
	return &jstore, nil
}

func (j JsonStore) GetRedirectByPath(path string) (domain.Redirect, error) {

	for i := 0; i < len(j.store); i++ {
		if strings.Compare(j.store[i].Path, path) == 0 {
			return j.store[i], nil
		}
	}
	return domain.Redirect{}, errors.New("Redirect not found")

}
