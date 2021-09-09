package service

import (
	"yn/todo/internal/model"
	"yn/todo/internal/payload"
	"yn/todo/internal/repository"
)

// Field
type TodoListService struct {
	repo repository.TodoList
}

// Constructor
func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

// Method
func (s *TodoListService) Create(userId int, list model.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}
func (s *TodoListService) GetAll(userId int) ([]model.TodoList, error) {
	return s.repo.GetAll(userId)
}
func (s *TodoListService) GetById(userId, listId int) (model.TodoList, error) {
	return s.repo.GetById(userId, listId)
}
func (s *TodoListService) DeleteById(userId, listId int) error {
	return s.repo.DeleteById(userId, listId)
}
func (s *TodoListService) UpdateById(userId, listId int, input payload.UpdateTodoListRequest) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateById(userId, listId, input)
}
func (s *TodoListService) GetAllItemsById(userId, listId int) ([]model.TodoItem, error) {
	return s.repo.GetAllItemsById(userId, listId)
}
