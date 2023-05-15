package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	ErrTitleCannotBeBlank       = errors.New("title can not be blank")
	ErrItemNotFound             = errors.New("item not found")
	ErrCannotUpdateFinishedItem = errors.New("can not update finished item")
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var statuses = [3]string{"Doing", "Done", "Deleted"}

func (item ItemStatus) String() string {
	return statuses[item]
}

func parseStr2ItemStatus(s string) (ItemStatus, error) {
	for i := range statuses {
		if statuses[i] == s {
			return ItemStatus(i), nil
		}
	}
	return ItemStatus(0), errors.New("invalid status string")
}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}
	strVal := string(bytes)
	v, err := parseStr2ItemStatus(strVal)
	if err != nil {
		return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
	}
	*item = v
	return nil
}

func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}

func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	val, err := parseStr2ItemStatus(str)
	if err != nil {
		return err
	}
	*item = val
	return nil
}

type SQLModel struct {
	Id        int        `json:"id,omitempty" gorm:"column:id"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

type TodoItem struct {
	SQLModel
	Title       string      `json:"title,omitempty" gorm:"column:title"`
	Description string      `json:"description,omitempty" gorm:"column:description"`
	Status      *ItemStatus `json:"status,omitempty" gorm:"column:status"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

func (item TodoItem) Validate() error {
	if item.Title == "" {
		return ErrTitleCannotBeBlank
	}
	return nil
}

type Paging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 || p.Limit >= 100 {
		p.Limit = 10
	}
}
