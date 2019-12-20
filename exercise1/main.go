package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/krishnakumarkp/gophercises/exercise1/csvstore"
	"github.com/krishnakumarkp/gophercises/exercise1/domain"
)

type Problem struct {
	store domain.ProblemStore
}

func (p Problem) GetAll() ([]domain.Problem, error) {
	problems, err := p.store.GetProblems()
	return problems, err
}

type Quiz struct {
	Problems []domain.Problem
	Correct  int
	Wrong    int
	Total    int
}

func (q *Quiz) AskQuestion() {
	areader := bufio.NewReader(os.Stdin)
	for _, p := range q.Problems {
		q.Total++
		fmt.Println(q.Total, ") ", p.Question)
		ans, _ := areader.ReadString('\n')
		ans = strings.Replace(ans, "\n", "", -1)

		if strings.Compare(ans, p.Answer) == 0 {
			q.Correct++
		} else {
			q.Wrong++
		}

	}
}

func main() {

	var fileName *string
	var qtime *int

	fileName = flag.String("file", "problems.csv", "Problems file")
	qtime = flag.Int("time", 30, "Test time")
	flag.Parse()

	problem := Problem{
		store: csvstore.NewCsvStore(*fileName),
	}

	problems, err := problem.GetAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	quiz := Quiz{problems, 0, 0, 0}

	ch := make(chan bool, 1)
	defer close(ch)

	go func() {
		quiz.AskQuestion()
		ch <- true
	}()

	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*qtime) * time.Second):
		fmt.Println("Quiz Timed out")

	}
	fmt.Printf("You answered %d questions out of which %d were correct and %d were incorrect!\n", quiz.Total, quiz.Correct, quiz.Wrong)

}
