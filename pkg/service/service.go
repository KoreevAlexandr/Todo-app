package service

import (
	"todo-app"
	"todo-app/pkg/repository"
)

type Authorisation interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, listId int) (todo.TodoList, error)
}

type TodoIdea interface {
}

type Service struct {
	Authorisation
	TodoList
	TodoIdea
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repos.Authorisation),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
