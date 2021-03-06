package general

import "github.com/jinzhu/gorm"

// GORMRepository .
type GORMRepository interface {
	DB() *gorm.DB
}

// QueryBuilder .
type QueryBuilder interface {
	Execute(db *gorm.DB, object interface{}) (e error)
	Order() interface{}
}
