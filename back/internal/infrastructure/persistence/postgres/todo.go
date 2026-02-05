package postgres

import (
	"github.com/bearl27/todos/back/internal/domain/model"
	"github.com/bearl27/todos/back/internal/domain/repository"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) Create(todo *model.Todo) error {
	return r.db.Create(todo).Error
}

func (r *todoRepository) FindAll() ([]*model.Todo, error) {
	var todos []*model.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *todoRepository) FindByID(id uint) (*model.Todo, error) {
	var todo model.Todo
	err := r.db.First(&todo, id).Error
	return &todo, err
}

func (r *todoRepository) Update(todo *model.Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) Delete(id uint) error {
	return r.db.Delete(&model.Todo{}, id).Error
}
