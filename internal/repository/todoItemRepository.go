package repository

import (
	"fmt"
	"strings"
	"yn/todo/internal/model"
	"yn/todo/internal/payload"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// Field
type TodoItemRepository struct {
	db *sqlx.DB
}

// Constructor
func NewTodoItemRepository(db *sqlx.DB) *TodoItemRepository {
	return &TodoItemRepository{db: db}
}

// Method
func (r *TodoItemRepository) Create(userId int, input payload.TodoItemRequest) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}

	var listId int
	if input.ListName != "" {
		query := fmt.Sprintf(`SELECT id FROM %s WHERE user_id=$1 AND name=$2`, todoListsTable)
		err = r.db.Get(&listId, query, userId, input.ListName)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	var id int
	createListQuery := fmt.Sprintf(`INSERT INTO %s (title, description, list_id, user_id) VALUES ($1, $2, $3, $4) RETURNING id`, todoItemsTable)
	row := tx.QueryRow(createListQuery, input.Title, input.Description, listId, userId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (r *TodoItemRepository) GetAll(userId int) ([]model.TodoItem, error) {
	var items []model.TodoItem
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id=$1`, todoItemsTable)
	if err := r.db.Select(&items, query, userId); err != nil {
		return nil, err
	}
	return items, nil
}
func (r *TodoItemRepository) GetById(userId, itemId int) (model.TodoItem, error) {
	var item model.TodoItem
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id=$1 AND id=$2`, todoItemsTable)
	err := r.db.Get(&item, query, userId, itemId)
	return item, err
}
func (r *TodoItemRepository) DeleteById(userId, itemId int) error {
	deleteItemsQuery := fmt.Sprintf(`DELETE FROM %s WHERE user_id=$1 AND id=$2`, todoItemsTable)
	_, err := r.db.Exec(deleteItemsQuery, userId, itemId)
	return err
}
func (r *TodoItemRepository) UpdateById(userId, itemId int, input payload.UpdateTodoItemRequest) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE user_id=$%d AND id=$%d",
		todoItemsTable, setQuery, argId, argId+1,
	)
	args = append(args, userId, itemId)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)
	_, err := r.db.Exec(query, args...)
	return err
}
