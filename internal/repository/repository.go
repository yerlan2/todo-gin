package repository

import (
	"yn/todo/internal/model"
	"yn/todo/internal/payload"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	usersTable     = "users"
	todoListsTable = "todo_lists"
	todoItemsTable = "todo_items"
)

// Interface
type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
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
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// Constructor
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		TodoList:      NewTodoListRepository(db),
		TodoItem:      NewTodoItemRepository(db),
	}
}
