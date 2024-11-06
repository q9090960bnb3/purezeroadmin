package svc

import (
	"backend/user-api/internal/config"
	"backend/user-api/models"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	TbUserModel models.TbUserModel
	TbRoleModel models.TbRoleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Dsn)
	return &ServiceContext{
		Config:      c,
		TbUserModel: models.NewTbUserModel(conn),
		TbRoleModel: models.NewTbRoleModel(conn),
	}
}
