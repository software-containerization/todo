package storage

import (
	"todo-api/cmd/config"
	"todo-api/pkg/data"

	_ "github.com/go-sql-driver/mysql" //MySQL driver
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// ItemRepository handles database interactions for todo items
type ItemRepository struct {
	db *gorm.DB
}

var Set = wire.NewSet(
	NewItemRepository,
	NewGormDB,
)

// NewGormDB provides the Gorm Database
func NewGormDB(c *config.Configuration) *gorm.DB {
	db, err := gorm.Open("mysql", c.DSN)
	if err != nil {
		logrus.Fatal(err)
	}

	db.AutoMigrate(&data.Item{})

	return db
}

// NewItemRepository returns a new ItemRepository
func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (i *ItemRepository) Create(item data.Item) error {
	return i.db.Create(&item).Error
}

func (i *ItemRepository) Update(item data.Item) error {
	return i.db.Save(&item).Error
}

func (i *ItemRepository) FindAll() ([]data.Item, error) {
	var items []data.Item
	result := i.db.Find(&items)
	return items, result.Error
}

func (i *ItemRepository) Delete(id string) error {
	return i.db.Delete(&data.Item{}, id).Error
}
