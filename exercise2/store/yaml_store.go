package store

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/krishnakumarkp/gophercises/exercise2/domain"
	"gopkg.in/yaml.v2"
)

type YamlStore struct {
	store []domain.Redirect
}

func NewYamlStore(file string) (*YamlStore, error) {
	var ystore YamlStore
	yamlFile, err := os.Open(file)
	if err != nil {
		return &ystore, err
	}
	byteValue, err := ioutil.ReadAll(yamlFile)
	if err != nil {
		return &ystore, err
	}
	err = yaml.Unmarshal(byteValue, &ystore.store)
	if err != nil {
		return &ystore, err
	}
	return &ystore, nil
}

func (y YamlStore) GetRedirectByPath(path string) (domain.Redirect, error) {

	for i := 0; i < len(y.store); i++ {
		if strings.Compare(y.store[i].Path, path) == 0 {
			return y.store[i], nil
		}
	}
	return domain.Redirect{}, errors.New("Redirect not found")

}
