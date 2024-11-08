package svc

import (
	"backend/user-api/internal/config"
	"backend/user-api/models"

	sqladapter "github.com/Blank-Xu/sql-adapter"
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
	db, err := conn.RawDB()
	if err != nil {
		logx.Errorf("Error on conn.RawDB: %+v", err)
		return nil
	}
	policy, err := sqladapter.NewAdapter(db, "mysql", "casbin_rule")
	if err != nil {
		logx.Errorf("Error on sqladapter.NewAdapter: %+v", err)
		return nil
	}
	enforcer, err := casbin.NewSyncedEnforcer(c.Casbin.ModelFile, policy)
	if err != nil {
		logx.Errorf("Error on casbin.NewSyncedEnforcer: %+v", err)
		return nil
	}
	// enforcer.AddPolicy("10086", "/10086", "get")
	return &ServiceContext{
		Config:        c,
		TbUserModel:   models.NewTbUserModel(conn),
		TbRoleModel:   models.NewTbRoleModel(conn),
		TbRouterModel: models.NewTbRouterModel(conn),
		Enforcer:      enforcer,
	}
}
