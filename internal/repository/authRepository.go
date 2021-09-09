package repository

import (
	"fmt"
	"yn/todo/internal/model"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Field
type AuthRepository struct {
	db *sqlx.DB
}

// Constructor
func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// Method
func (r *AuthRepository) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (username, password_hash) VALUES ($1, $2) RETURNING id`, usersTable)
	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (r *AuthRepository) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf(`SELECT id FROM %s WHERE username=$1 AND password_hash=$2`, usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
