// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tbRouterFieldNames          = builder.RawFieldNames(&TbRouter{})
	tbRouterRows                = strings.Join(tbRouterFieldNames, ",")
	tbRouterRowsExpectAutoSet   = strings.Join(stringx.Remove(tbRouterFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tbRouterRowsWithPlaceHolder = strings.Join(stringx.Remove(tbRouterFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tbRouterModel interface {
		Insert(ctx context.Context, data *TbRouter) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TbRouter, error)
		Update(ctx context.Context, data *TbRouter) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTbRouterModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TbRouter struct {
		Id        int64          `db:"id"`
		ParentId  int64          `db:"parent_id"`
		MenuType  int64          `db:"menu_type"`
		Path      string         `db:"path"`
		Name      string         `db:"name"`
		Component string         `db:"component"`
		MetaTitle string         `db:"meta_title"`
		MetaIcon  string         `db:"meta_icon"`
		MetaRank  int64          `db:"meta_rank"`
		MetaRoles sql.NullString `db:"meta_roles"`
		MetaAuths sql.NullString `db:"meta_auths"`
		CreateTs  int64          `db:"create_ts"`
		UpdateTs  int64          `db:"update_ts"`
	}
)

func newTbRouterModel(conn sqlx.SqlConn) *defaultTbRouterModel {
	return &defaultTbRouterModel{
		conn:  conn,
		table: "`tb_router`",
	}
}

func (m *defaultTbRouterModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTbRouterModel) FindOne(ctx context.Context, id int64) (*TbRouter, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tbRouterRows, m.table)
	var resp TbRouter
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTbRouterModel) Insert(ctx context.Context, data *TbRouter) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, tbRouterRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ParentId, data.MenuType, data.Path, data.Name, data.Component, data.MetaTitle, data.MetaIcon, data.MetaRank, data.MetaRoles, data.MetaAuths, data.CreateTs, data.UpdateTs)
	return ret, err
}

func (m *defaultTbRouterModel) Update(ctx context.Context, data *TbRouter) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tbRouterRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ParentId, data.MenuType, data.Path, data.Name, data.Component, data.MetaTitle, data.MetaIcon, data.MetaRank, data.MetaRoles, data.MetaAuths, data.CreateTs, data.UpdateTs, data.Id)
	return err
}

func (m *defaultTbRouterModel) tableName() string {
	return m.table
}
