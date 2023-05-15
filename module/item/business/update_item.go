package business

import (
	"context"

	"github.com/tranlequocthong313/todolist/module/item/model"
)

type UpdateTodoItemStorage interface {
	FindItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, condition map[string]interface{}, dataUpdate *model.TodoItem) error
}

type updateBusiness struct {
	store UpdateTodoItemStorage
}

func NewUpdateTodoItemBusiness(store UpdateTodoItemStorage) *updateBusiness {
	return &updateBusiness{store}
}

func (business *updateBusiness) UpdateItem(ctx context.Context, condition map[string]interface{}, dataUpdate *model.TodoItem) error {
	oldItem, err := business.store.FindItem(ctx, condition)
	if err != nil {
		return err
	}
	if oldItem.Status.String() == "Deleted" {
		return model.ErrCannotUpdateFinishedItem
	}
	if err := business.store.UpdateItem(ctx, condition, dataUpdate); err != nil {
		return err
	}
	return nil
}
