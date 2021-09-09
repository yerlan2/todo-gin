package model

type TodoList struct {
	Id     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name" binding:"required"`
	UserId int    `json:"user_id" db:"user_id"`
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
	ListId      int    `json:"list_id" db:"list_id"`
	UserId      int    `json:"user_id" db:"user_id"`
}
