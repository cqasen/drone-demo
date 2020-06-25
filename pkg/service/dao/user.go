package dao

import (
	"github.com/cqasen/gin-demo/pkg/service/entity"
	"github.com/jinzhu/gorm"
)

type UserDao struct {
	Dao
}

func User(db *gorm.DB) *UserDao {
	return &UserDao{Dao{db: db}}
}

func (dao UserDao) Get(name string) (*entity.ZbpMember, error) {
	member := new(entity.ZbpMember)
	query := dao.db.Table(entity.TABLE_MEMBER).Where(entity.ZbpMember{
		MemName: name,
	}).Debug().First(&member)
	if err := query.Error; err != nil {
		return nil, err
	}
	return member, nil
}
