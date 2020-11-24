package dao

import "gorm.io/gorm"

type Dao struct {
	db *gorm.DB
}
