package service

import (
	"yn/todo/internal/model"
	"yn/todo/internal/payload"
	"yn/todo/internal/repository"
)

// Field
type TodoItemService struct {
	repo repository.TodoItem
}

// Constructor
func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}

// Method
func (s *TodoItemService) Create(userId int, input payload.TodoItemRequest) (int, error) {
	return s.repo.Create(userId, input)
}
func (s *TodoItemService) GetAll(userId int) ([]model.TodoItem, error) {
	return s.repo.GetAll(userId)
}
func (s *TodoItemService) GetById(userId, itemId int) (model.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}
func (s *TodoItemService) DeleteById(userId, itemId int) error {
	return s.repo.DeleteById(userId, itemId)
}
func (s *TodoItemService) UpdateById(userId, itemId int, input payload.UpdateTodoItemRequest) error {
	return s.repo.UpdateById(userId, itemId, input)
}
