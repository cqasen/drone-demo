package dao

import (
	"github.com/cqasen/gin-demo/pkg/service/entity"
	"gorm.io/gorm"
)

type UserDao struct {
	Dao
}

func User(db *gorm.DB) *UserDao {
	return &UserDao{Dao{db: db}}
}

func (dao UserDao) Get(name string) (*entity.ZbpMember, error) {
	member := new(entity.ZbpMember)
	query := dao.db.Table(entity.TableMember).Where(entity.ZbpMember{
		MemName: name,
	}).Debug().First(&member)
	if err := query.Error; err != nil {
		return nil, err
	}
	return member, nil
}
