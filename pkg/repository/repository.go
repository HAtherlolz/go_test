package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/Hatherlolz/go_test"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}


type TodoList interface {
	
}


type TodoItem interface {
	
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
