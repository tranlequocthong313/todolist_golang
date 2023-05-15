package storage

import (
	"context"

	"github.com/tranlequocthong313/todolist/module/item/model"
)

func (s *mysqlStorage) CreateItem(ctx context.Context, data *model.TodoItem) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
