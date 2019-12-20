package csvstore

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"os"

	"github.com/krishnakumarkp/gophercises/exercise1/domain"
)

type CsvStore struct {
	file string
}

func NewCsvStore(file string) *CsvStore {
	return &CsvStore{file}
}

func (c CsvStore) GetProblems() ([]domain.Problem, error) {

	var problems []domain.Problem
	var csvFile *os.File
	var err error

	if csvFile, err = os.Open(c.file); err != nil {
		return problems, errors.New("Could not open the csv file")
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}
		if error != nil {
			return problems, errors.New("Could not read the csv file")
		}
		p := domain.Problem{line[0], line[1]}
		problems = append(problems, p)
	}
	return problems, nil
}
