package models

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbRoleModel = (*customTbRoleModel)(nil)

type (
	// TbRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbRoleModel.
	TbRoleModel interface {
		tbRoleModel
		withSession(session sqlx.Session) TbRoleModel
		FindAll(ctx context.Context) ([]*TbRole, error)
	}

	customTbRoleModel struct {
		*defaultTbRoleModel
	}
)

// NewTbRoleModel returns a model for the database table.
func NewTbRoleModel(conn sqlx.SqlConn) TbRoleModel {
	return &customTbRoleModel{
		defaultTbRoleModel: newTbRoleModel(conn),
	}
}

func (m *customTbRoleModel) withSession(session sqlx.Session) TbRoleModel {
	return NewTbRoleModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTbRoleModel) FindAll(ctx context.Context) ([]*TbRole, error) {
	var list []*TbRole
	err := m.conn.QueryRowsCtx(ctx, &list, "select * from tb_role")
	return list, err
}
