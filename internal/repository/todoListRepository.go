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
type TodoListRepository struct {
	db *sqlx.DB
}

// Constructor
func NewTodoListRepository(db *sqlx.DB) *TodoListRepository {
	return &TodoListRepository{db: db}
}

// Method
func (r *TodoListRepository) Create(userId int, list model.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, nil
	}
	var id int
	createListQuery := fmt.Sprintf(`INSERT INTO %s (name, user_id) VALUES ($1, $2) RETURNING id`, todoListsTable)
	row := tx.QueryRow(createListQuery, list.Name, userId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (r *TodoListRepository) GetAll(userId int) ([]model.TodoList, error) {
	var lists []model.TodoList
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id=$1`, todoListsTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}
func (r *TodoListRepository) GetById(userId, listId int) (model.TodoList, error) {
	var list model.TodoList
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id=$1 AND id=$2`, todoListsTable)
	err := r.db.Get(&list, query, userId, listId)
	return list, err
}
func (r *TodoListRepository) DeleteById(userId, listId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	deleteItemsQuery := fmt.Sprintf(`DELETE FROM %s WHERE user_id=$1 AND list_id=$2`, todoItemsTable)
	_, err = tx.Exec(deleteItemsQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return err
	}
	deleteListQuery := fmt.Sprintf(`DELETE FROM %s WHERE user_id=$1 AND id=$2`, todoListsTable)
	_, err = tx.Exec(deleteListQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func (r *TodoListRepository) UpdateById(userId, listId int, input payload.UpdateTodoListRequest) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE  user_id=$%d AND id=$%d",
		todoListsTable, setQuery, argId, argId+1,
	)
	args = append(args, userId, listId)
	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)
	_, err := r.db.Exec(query, args...)
	return err
}
func (r *TodoListRepository) GetAllItemsById(userId, listId int) ([]model.TodoItem, error) {
	var items []model.TodoItem
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id=$1 AND list_id=$2`, todoItemsTable)
	if err := r.db.Select(&items, query, userId, listId); err != nil {
		return nil, err
	}
	return items, nil
}
