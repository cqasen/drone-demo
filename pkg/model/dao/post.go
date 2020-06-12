package dao

import (
	"fmt"
	"github.com/cqasen/gin-demo/pkg/model/entity"
	"github.com/jinzhu/gorm"
)

type PostDao struct {
	Dao
}

func Post(db *gorm.DB) *PostDao {
	return &PostDao{Dao{db: db}}
}

func (dao PostDao) Get(id int32) (*entity.ZbpPost, error) {
	post := new(entity.ZbpPost)
	query := dao.db.Table(entity.TABLE_POST).Debug().Where(entity.ZbpPost{LogID: id}).First(post)
	if err := query.Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (dao PostDao) GetList() ([]entity.ZbpPost, error) {
	var post []entity.ZbpPost
	query := dao.db.Table(entity.TABLE_POST).Debug().Order("log_ID DESC").Scan(&post)
	if err := query.Error; err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return post, nil
}
