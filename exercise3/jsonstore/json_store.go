package jsonstore

import (
	"encoding/json"
	"errors"
	"github.com/krishnakumarkp/gophercises/exercise3/domain"
	"io/ioutil"
	"os"
)

type JsonStore struct {
	store map[string]domain.Story
}

func NewStore(file string) (*JsonStore, error) {
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

func (j JsonStore) GetByArc(arc string) (domain.Story, error) {
	if v, ok := j.store[arc]; ok {
		return v, nil
	}
	return domain.Story{}, errors.New("Story not found")

}
