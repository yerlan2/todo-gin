package service

import (
	"yn/todo/internal/model"
	"yn/todo/internal/payload"
	"yn/todo/internal/repository"
)

// Interface
type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {
	Create(userId int, list model.TodoList) (int, error)
	GetAll(userId int) ([]model.TodoList, error)
	GetById(userId, listId int) (model.TodoList, error)
	DeleteById(userId, listId int) error
	UpdateById(userId, listId int, input payload.UpdateTodoListRequest) error
	GetAllItemsById(userId, listId int) ([]model.TodoItem, error)
}
type TodoItem interface {
	Create(userId int, input payload.TodoItemRequest) (int, error)
	GetAll(userId int) ([]model.TodoItem, error)
	GetById(userId, itemId int) (model.TodoItem, error)
	DeleteById(userId, itemId int) error
	UpdateById(userId, itemId int, input payload.UpdateTodoItemRequest) error
}

// Field
type Service struct {
	Authorization
	TodoList
	TodoItem
}

// Constructor
func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:      NewTodoItemService(repo.TodoItem),
	}
}
