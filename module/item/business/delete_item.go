package business

import (
	"context"

	"github.com/tranlequocthong313/todolist/module/item/model"
)

type DeleteTodoItemStorage interface {
	FindItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, condition map[string]interface{}) error
}

type deleteBusinuess struct {
	store DeleteTodoItemStorage
}

func NewDeleteTodoItemBusiness(store DeleteTodoItemStorage) *deleteBusinuess {
	return &deleteBusinuess{store}
}

func (business *deleteBusinuess) DeleteItem(ctx context.Context, condition map[string]interface{}) error {
	_, err := business.store.FindItem(ctx, condition)
	if err != nil {
		return err
	}
	if err := business.store.DeleteItem(ctx, condition); err != nil {
		return err
	}
	return nil
}
