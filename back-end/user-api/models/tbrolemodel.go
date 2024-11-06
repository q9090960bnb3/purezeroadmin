package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TbRoleModel = (*customTbRoleModel)(nil)

type (
	// TbRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbRoleModel.
	TbRoleModel interface {
		tbRoleModel
		withSession(session sqlx.Session) TbRoleModel
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
