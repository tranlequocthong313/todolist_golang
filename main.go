package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/tranlequocthong313/todolist/module/item/transport"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	connStr, found := os.LookupEnv("DB_CONN")
	if !found {
		log.Fatalln("Missing MySQL connection string.")
	}

	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/items", transport.HandleListItems(db))
		v1.GET("/items/:id", transport.HandleFindAnItem(db))
		v1.POST("/items", transport.HandleCreateItem(db))
		v1.PUT("/items/:id", transport.HandleUpdateAnItem(db))
		v1.DELETE("/items/:id", transport.HandleDeleteAnItem(db))
	}
	router.Run(":3000")
}

// import (
//     "database/sql/driver"
//     "errors"
//     "fmt"
//     "log"
//     "net/http"
//     "os"
//     "strconv"
//     "strings"
//
//     "github.com/gin-gonic/gin"
//     "github.com/tranlequocthong313/todolist/common"
//     "gorm.io/driver/mysql"
//     "gorm.io/gorm"
// )
//
// type ItemStatus int
//
// const (
//     ItemStatusDoing ItemStatus = iota
//     ItemStatusDone
//     ItemStatusDeleted
// )
//
// var statuses = [3]string{"Doing", "Done", "Deleted"}
//
// func (item ItemStatus) String() string {
//     return statuses[item]
// }
//
// func parseStr2ItemStatus(s string) (ItemStatus, error) {
//     for i := range statuses {
//         if statuses[i] == s {
//             return ItemStatus(i), nil
//         }
//     }
//     return ItemStatus(0), errors.New("invalid status string")
// }
//
// func (item *ItemStatus) Scan(value interface{}) error {
//     bytes, ok := value.([]byte)
//     if !ok {
//         return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
//     }
//     strVal := string(bytes)
//     v, err := parseStr2ItemStatus(strVal)
//     if err != nil {
//         return errors.New(fmt.Sprintf("fail to scan data from sql: %s", value))
//     }
//     *item = v
//     return nil
// }
//
// func (item *ItemStatus) Value() (driver.Value, error) {
//     if item == nil {
//         return nil, nil
//     }
//     return item.String(), nil
// }
//
// func (item *ItemStatus) MarshalJSON() ([]byte, error) {
//     if item == nil {
//         return nil, nil
//     }
//     return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
// }
//
// func (item *ItemStatus) UnmarshalJSON(data []byte) error {
//     str := strings.ReplaceAll(string(data), "\"", "")
//     val, err := parseStr2ItemStatus(str)
//     if err != nil {
//         return err
//     }
//     *item = val
//     return nil
// }
//
// type TodoItem struct {
//     common.SQLModel
//     Title       string      `json:"title,omitempty" gorm:"column:title"`
//     Description string      `json:"description,omitempty" gorm:"column:description"`
//     Status      *ItemStatus `json:"status,omitempty" gorm:"column:status"`
// }
//
// func (TodoItem) TableName() string {
//     return "todo_items"
// }
//
// type TodoItemCreation struct {
//     Id          int         `json:"-" gorm:"column:id;"`
//     Title       string      `json:"title" gorm:"column:title;"`
//     Description string      `json:"description" gorm:"column:description;"`
//     Status      *ItemStatus `json:"status" gorm:"column:status;"`
// }
//
// func (TodoItemCreation) TableName() string {
//     return TodoItem{}.TableName()
// }
//
// type TodoItemUpdate struct {
//     Title       *string     `json:"title" gorm:"column:title;"`
//     Description *string     `json:"description" gorm:"column:description;"`
//     Status      *ItemStatus `json:"status" gorm:"column:status;"`
// }
//
// func (TodoItemUpdate) TableName() string {
//     return TodoItem{}.TableName()
// }
//
// func main() {
//     dbConn, ok := os.LookupEnv("DB_CONN")
//     if !ok {
//         log.Fatalln("Missing MySQL connection string")
//     }
//     db, err := gorm.Open(mysql.Open(dbConn), &gorm.Config{})
//     if err != nil {
//         log.Fatalln("Cannot connect to MySQL:", err)
//     }
//     fmt.Println("Connected to MYSQL", db)
//
//     r := gin.Default()
//
//     v1 := r.Group("/v1")
//     {
//         items := v1.Group("/items")
//         {
//             items.POST("", CreateItem(db))
//             items.GET("", GetItems(db))
//             items.GET("/:id", GetItem(db))
//             items.PUT("/:id", UpdateItem(db))
//             items.DELETE("/:id", DeleteItem(db))
//         }
//     }
//     r.GET("/ping", func(c *gin.Context) {
//         c.JSON(http.StatusOK, gin.H{
//             "message": "pong",
//         })
//     })
//
//     r.Run(":3000")
// }
//
// func CreateItem(db *gorm.DB) func(*gin.Context) {
//     return func(c *gin.Context) {
//         var todo TodoItemCreation
//         if err := c.ShouldBind(&todo); err != nil {
//             c.JSON(http.StatusBadRequest, gin.H{
//                 "error": err.Error(),
//             })
//             return
//         }
//         if err := db.Create(&todo).Error; err != nil {
//             c.JSON(http.StatusBadRequest, gin.H{
//                 "error": err.Error(),
//             })
//             return
//         }
//         c.JSON(http.StatusOK, common.SimpleSuccessResponse(todo.Id))
//     }
// }
// func GetItems(db *gorm.DB) func(*gin.Context) {
//     return func(c *gin.Context) {
//         var paging common.Paging
//         if err := c.ShouldBind(&paging); err != nil {
//             c.JSON(http.StatusBadRequest, gin.H{
//                 "error": err.Error(),
//             })
//             return
//         }
//         paging.Process()
//         var todos []TodoItem
//         db = db.Where("status <> ?", "Deleted")
//         if err := db.Table(TodoItem{}.TableName()).Count(&paging.Total).Error; err != nil {
//             c.JSON(http.StatusBadRequest, gin.H{
//                 "error": err.Error(),
//             })
//             return
//         }
//         if err := db.Order("id desc").
//             Offset((paging.Page - 1) * paging.Limit).
//             Limit(paging.Limit).
//             Find(&todos).Error; err != nil {
//             c.JSON(http.StatusBadRequest, gin.H{
//                 "error": err.Error(),
//             })
//             return
//         }
//         c.JSON(http.StatusOK, common.NewSuccessResponse(todos, paging, nil))
//     }
// }
// func GetItem(db *gorm.DB) func(*gin.Context) {
//     return func(c *gin.Context) {
//         var todo TodoItem
//         id, err := strconv.Atoi(c.Param("id"))
//         if err != nil {
//             c.JSON(http.StatusBadRequest, gin.H{
//                 "error": err.Error(),
//             })
//             return
//         }
//         if err := db.Where("id = ?", id).First(&todo).Error; err != nil {
//             c.JSON(http.StatusBadRequest, gin.H{
//                 "error": err.Error(),
//             })
//             return
//         }
//         c.JSON(http.StatusOK, common.SimpleSuccessResponse(todo))
//     }
// }
// func UpdateItem(db *gorm.DB) func(*gin.Context) {
//     return func(c *gin.Context) {
//         var todo TodoItemUpdate
//         id, err := strconv.Atoi(c.Param("id"))
//         if err != nil {
//             c.JSON(http.StatusBadRequest, gin.H{
//                 "error": err.Error(),
//             })
//             return
//         }
//         if err := c.ShouldBind(&todo); err != nil {
//             c.JSON(http.StatusBadRequest, gin.H{
//                 "error": err.Error(),
//             })
//             return
//         }
//         if err := db.Where("id = ?", id).Updates(&todo).Error; err != nil {
//             c.JSON(http.StatusBadRequest, gin.H{
//                 "error": err.Error(),
//             })
//             return
//         }
//         c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
//     }
// }
// func DeleteItem(db *gorm.DB) func(*gin.Context) {
//     return func(c *gin.Context) {
//         id, err := strconv.Atoi(c.Param("id"))
//         if err != nil {
//             c.JSON(http.StatusBadRequest, gin.H{
//                 "error": err.Error(),
//             })
//             return
//         }
//         if err := db.Table(TodoItem{}.TableName()).
//             Where("id = ?", id).
//             Updates(map[string]interface{}{
//                 "status": "Deleted",
//             }).Error; err != nil {
//             c.JSON(http.StatusBadRequest, gin.H{
//                 "error": err.Error(),
//             })
//             return
//         }
//         c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
//     }
// }
