package service

import "todo-app/pkg/repository"

type Authorisation interface {
}

type TodoList interface {
}

type TodoIdea interface {
}

type Service struct {
	Authorisation
	TodoList
	TodoIdea
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
