package storage

import (
	"context"

	"github.com/tranlequocthong313/todolist/module/item/model"
)

func (s *mysqlStorage) UpdateItem(ctx context.Context, condition map[string]interface{}, dataUpdate *model.TodoItem) error {
	if err := s.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}
	return nil
}
