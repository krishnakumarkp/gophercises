package domain

type Problem struct {
	Question string
	Answer   string
}

type ProblemStore interface {
	GetProblems() ([]Problem, error)
}
