package svc

import (
	"backend/user-api/internal/config"
	"backend/user-api/models"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	TbUserModel   models.TbUserModel
	TbRoleModel   models.TbRoleModel
	TbRouterModel models.TbRouterModel
	Enforcer      *casbin.SyncedEnforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Dsn)
	enforcer, err := casbin.NewSyncedEnforcer(c.Casbin.ModelFile, c.Casbin.PolicyFile)
	if err != nil {
		logx.Errorf("Error on NewSyncedEnforcer: %+v", err)
		return nil
	}
	return &ServiceContext{
		Config:        c,
		TbUserModel:   models.NewTbUserModel(conn),
		TbRoleModel:   models.NewTbRoleModel(conn),
		TbRouterModel: models.NewTbRouterModel(conn),
		Enforcer:      enforcer,
	}
}
