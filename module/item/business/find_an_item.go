package business

import (
	"context"

	"github.com/tranlequocthong313/todolist/module/item/model"
)

type FindTodoItemStorage interface {
	FindItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error)
}

type findBusiness struct {
	store FindTodoItemStorage
}

func NewFindTodoItemBusiness(store FindTodoItemStorage) *findBusiness {
	return &findBusiness{store}
}

func (business *findBusiness) FindAnItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error) {
	item, err := business.store.FindItem(ctx, condition)
	if err != nil {
		return nil, err
	}
	return item, nil
}
