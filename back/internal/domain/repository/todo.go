package repository

import "github.com/bearl27/todos/back/internal/domain/model"

type TodoRepository interface {
	Create(todo *model.Todo) error
	FindAll() ([]*model.Todo, error)
	FindByID(id uint) (*model.Todo, error)
	Update(todo *model.Todo) error
	Delete(id uint) error
}
