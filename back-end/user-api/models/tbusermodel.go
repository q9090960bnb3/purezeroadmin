package models

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TbUserModel = (*customTbUserModel)(nil)

type (
	// TbUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbUserModel.
	TbUserModel interface {
		tbUserModel
		withSession(session sqlx.Session) TbUserModel
	}

	customTbUserModel struct {
		*defaultTbUserModel
	}
)

// NewTbUserModel returns a model for the database table.
func NewTbUserModel(conn sqlx.SqlConn) TbUserModel {
	return &customTbUserModel{
		defaultTbUserModel: newTbUserModel(conn),
	}
}

func (m *customTbUserModel) withSession(session sqlx.Session) TbUserModel {
	return NewTbUserModel(sqlx.NewSqlConnFromSession(session))
}
