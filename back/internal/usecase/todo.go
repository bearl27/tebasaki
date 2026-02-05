package usecase

import (
	"github.com/bearl27/todos/back/internal/domain/model"
	"github.com/bearl27/todos/back/internal/domain/repository"
)

type TodoUsecase interface {
	CreateTodo(title string) error
	GetTodos() ([]*model.Todo, error)
	GetTodo(id uint) (*model.Todo, error)
	UpdateTodo(id uint, title string, completed bool) error
	DeleteTodo(id uint) error
}

type todoUsecase struct {
	repo repository.TodoRepository
}

func NewTodoUsecase(repo repository.TodoRepository) TodoUsecase {
	return &todoUsecase{repo: repo}
}

func (u *todoUsecase) CreateTodo(title string) error {
	todo := &model.Todo{Title: title, Completed: false}
	return u.repo.Create(todo)
}

func (u *todoUsecase) GetTodos() ([]*model.Todo, error) {
	return u.repo.FindAll()
}

func (u *todoUsecase) GetTodo(id uint) (*model.Todo, error) {
	return u.repo.FindByID(id)
}

func (u *todoUsecase) UpdateTodo(id uint, title string, completed bool) error {
	todo, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}
	todo.Title = title
	todo.Completed = completed
	return u.repo.Update(todo)
}

func (u *todoUsecase) DeleteTodo(id uint) error {
	return u.repo.Delete(id)
}
