package storage

import (
	"context"

	"github.com/tranlequocthong313/todolist/module/item/model"
	"gorm.io/gorm"
)

func (s *mysqlStorage) FindItem(ctx context.Context, condition map[string]interface{}) (*model.TodoItem, error) {
	var item model.TodoItem
	if err := s.db.Where(condition).First(&item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, model.ErrItemNotFound
		}
		return nil, err
	}
	return &item, nil
}
