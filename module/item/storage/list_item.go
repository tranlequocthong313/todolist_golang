package storage

import (
	"context"

	"github.com/tranlequocthong313/todolist/module/item/model"
)

func (s *mysqlStorage) ListItems(ctx context.Context, condition map[string]interface{}, paging *model.Paging) ([]model.TodoItem, error) {
	var todos []model.TodoItem
	offset := (paging.Page - 1) * paging.Limit
	if err := s.db.
		Table(model.TodoItem{}.TableName()).
		Where(condition).
		Count(&paging.Total).
		Offset(offset).
		Limit(paging.Limit).
		Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
