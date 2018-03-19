package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
)

//Quiz is an ordered series of questions
type Quiz struct {
	questions      []QuizQuestion
	correct, wrong int
}

type QuizQuestion struct {
	Subject  string
	Question string
	Answer   QuizAnswer
}

type QuizAnswer struct {
	Text   string
	Number int
	Units  string
}

type QuizBank struct {
	Questions []QuizQuestion
}

func load(path string) *QuizBank {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	q := new(QuizBank)
	err = json.Unmarshal(b, q)
	if err != nil {
		panic(err)
	}
	return q
}

type QuizPolicy = map[string]string

func newQuiz(bank *QuizBank, hints QuizPolicy) *Quiz {
	q := new(Quiz)

	switch hints["selection"] {
	case "all":
		q.questions = make([]QuizQuestion, len(bank.Questions))
		copy(q.questions, bank.Questions)
	}

	switch hints["ordering"] {
	case "shuffle":
		n := len(q.questions)
		shuffled := make([]QuizQuestion, n)
		r := rand.New(rand.NewSource(time.Now().Unix()))
		p := r.Perm(n)
		for i, v := range p {
			shuffled[i] = q.questions[v]
		}
		q.questions = shuffled
	}

	return q
}
