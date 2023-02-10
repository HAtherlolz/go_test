package service

import (
	"github.com/Hatherlolz/go_test/pkg/repository"
	"github.com/Hatherlolz/go_test"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}


type TodoList interface {
	
}


type TodoItem interface {
	
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}