package models

import (
	"context"
	"fmt"

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
		FindList(ctx context.Context, name, code, status string, page, pageSize int64) ([]*TbRole, int64, error)
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

func (m *customTbRoleModel) FindList(ctx context.Context, name, code, status string, page, pageSize int64) ([]*TbRole, int64, error) {
	base := ""
	needAnd := false
	if name != "" {
		base += fmt.Sprintf("name = '%s'", name)
		needAnd = true
	}
	if code != "" {
		if needAnd {
			base += " and "
		}
		base += fmt.Sprintf("code = '%s'", code)
	}
	if status != "" {
		if needAnd {
			base += " and "
		}
		base += fmt.Sprintf("status = %s", status)
	}
	if base != "" {
		base = " where " + base
	}

	var total int64
	err := m.conn.QueryRowCtx(ctx, &total, "select count(*) from tb_role"+base)
	if err != nil {
		return nil, 0, err
	}
	var list []*TbRole
	err = m.conn.QueryRowsCtx(ctx, &list, fmt.Sprintf("select * from tb_role%s limit ?,?", base), (page-1)*pageSize, pageSize)
	return list, total, err
}
