package storage

import (
	"context"

	"github.com/tranlequocthong313/todolist/module/item/model"
)

func (s *mysqlStorage) DeleteItem(ctx context.Context, condition map[string]interface{}) error {
	if err := s.db.Table(model.TodoItem{}.TableName()).Where(condition).Updates(map[string]interface{}{
		"status": "Deleted",
	}).Error; err != nil {
		return err
	}
	return nil
}
