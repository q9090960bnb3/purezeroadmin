package models

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TbRouterModel = (*customTbRouterModel)(nil)

type (
	// TbRouterModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbRouterModel.
	TbRouterModel interface {
		tbRouterModel
		withSession(session sqlx.Session) TbRouterModel
		FindAllFromParentID(ctx context.Context, parentID int64) ([]*TbRouter, error)
		FindAll(ctx context.Context) ([]*TbRouter, error)
	}

	customTbRouterModel struct {
		*defaultTbRouterModel
	}
)

// NewTbRouterModel returns a model for the database table.
func NewTbRouterModel(conn sqlx.SqlConn) TbRouterModel {
	return &customTbRouterModel{
		defaultTbRouterModel: newTbRouterModel(conn),
	}
}

func (m *customTbRouterModel) withSession(session sqlx.Session) TbRouterModel {
	return NewTbRouterModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customTbRouterModel) FindAllFromParentID(ctx context.Context, parentID int64) ([]*TbRouter, error) {
	var resp []*TbRouter
	err := m.conn.QueryRowsCtx(ctx, &resp, "select * from tb_router where parent_id = ?", parentID)
	return resp, err
}

func (m *customTbRouterModel) FindAll(ctx context.Context) ([]*TbRouter, error) {
	var resp []*TbRouter
	err := m.conn.QueryRowsCtx(ctx, &resp, "select * from tb_router")
	return resp, err
}
