package business

import (
	"context"

	"github.com/tranlequocthong313/todolist/module/item/model"
)

type CreateTodoItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItem) error
}

type createBusiness struct {
	store CreateTodoItemStorage
}

func NewCreateToDoItemBusiness(store CreateTodoItemStorage) *createBusiness {
	return &createBusiness{store}
}

func (business *createBusiness) CreateNewItem(ctx context.Context, data *model.TodoItem) error {
	if err := data.Validate(); err != nil {
		return err
	}
	if err := business.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}
