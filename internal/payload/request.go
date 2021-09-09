package payload

import "errors"

type SignInRequest struct {
	Username *string `json:"username" binding:"required"`
	Password *string `json:"password" binding:"required"`
}

type UpdateTodoListRequest struct {
	Name *string `json:"name"`
}

func (i UpdateTodoListRequest) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type TodoItemRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	ListName    string `json:"list_name"`
}

type UpdateTodoItemRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i UpdateTodoItemRequest) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
