package store

import (
	"errors"

	"github.com/krishnakumarkp/gophercises/exercise2/domain"
)

type MapStore struct {
	store map[string]domain.Redirect
}

func NewMapStore() *MapStore {
	mstore := &MapStore{store: make(map[string]domain.Redirect)}
	mstore.store["/dogs"] = domain.Redirect{"/dogs", "https://www.petfinder.com/dog-breeds/"}
	return mstore
}

func (m MapStore) GetRedirectByPath(path string) (domain.Redirect, error) {
	if v, ok := m.store[path]; ok {
		return v, nil
	}
	return domain.Redirect{}, errors.New("Redirect not found")

}
