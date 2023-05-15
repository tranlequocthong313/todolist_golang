package business

import (
	"context"

	"github.com/tranlequocthong313/todolist/module/item/model"
)

type ListTodoItemStorage interface {
	ListItems(ctx context.Context, condition map[string]interface{}, paging *model.Paging) ([]model.TodoItem, error)
}

type listBusiness struct {
	store ListTodoItemStorage
}

func NewListTodoItemBusiness(store ListTodoItemStorage) *listBusiness {
	return &listBusiness{store}
}

func (business *listBusiness) ListItems(ctx context.Context, condition map[string]interface{}, paging *model.Paging) ([]model.TodoItem, error) {
	items, err := business.store.ListItems(ctx, condition, paging)
	if err != nil {
		return nil, err
	}
	return items, nil
}
