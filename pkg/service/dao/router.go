package dao

import (
	"github.com/cqasen/gin-demo/pkg/service/entity"
	"github.com/jinzhu/gorm"
)

type RouterDao struct {
	Dao
}

func NewRouter(db *gorm.DB) *RouterDao {
	return &RouterDao{Dao{db: db}}
}

func (dao RouterDao) Save(route entity.Routers) entity.Routers {
	dao.db.Save(&route)
	return route
}

func (dao RouterDao) GetOne(path string, method string) entity.Routers {
	var route entity.Routers
	dao.db.Table(entity.TABLE_ROUTERS).
		Where(entity.Routers{Path: path, Method: method}).
		First(&route)
	return route
}
