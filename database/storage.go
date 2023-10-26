package database

import (
	"github.com/google/uuid"
)

type Account struct {
	id       uuid.UUID
	name     string
	password string
}

type Question struct {
	id      uuid.UUID
	title   string
	content string
	tags    []string
}

type Answer struct {
	id         uuid.UUID
	answer     string
	quetion_id uuid.UUID
}

type Storage interface {
	insertQuestion(question string) (Question, error)
	updateQuestion(question string) (Question, error)
	deleteQuestion(id uuid.UUID) error
	getQuestions() ([]Question, error)
	insertAnswer(answer, accountId uuid.UUID) (Answer, error)

	insertAccound(account string) error
	getAccount(email string) (Account, error)
}
