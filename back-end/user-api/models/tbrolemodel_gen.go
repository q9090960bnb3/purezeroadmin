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
	tbRoleFieldNames          = builder.RawFieldNames(&TbRole{})
	tbRoleRows                = strings.Join(tbRoleFieldNames, ",")
	tbRoleRowsExpectAutoSet   = strings.Join(stringx.Remove(tbRoleFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tbRoleRowsWithPlaceHolder = strings.Join(stringx.Remove(tbRoleFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tbRoleModel interface {
		Insert(ctx context.Context, data *TbRole) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TbRole, error)
		FindOneByCode(ctx context.Context, code string) (*TbRole, error)
		Update(ctx context.Context, data *TbRole) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTbRoleModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TbRole struct {
		Id          int64  `db:"id"`
		Code        string `db:"code"`
		Name        string `db:"name"`
		Status      int64  `db:"status"`
		Remark      string `db:"remark"`
		Permissions string `db:"permissions"`
		CreateTs    int64  `db:"create_ts"`
		UpdateTs    int64  `db:"update_ts"`
	}
)

func newTbRoleModel(conn sqlx.SqlConn) *defaultTbRoleModel {
	return &defaultTbRoleModel{
		conn:  conn,
		table: "`tb_role`",
	}
}

func (m *defaultTbRoleModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTbRoleModel) FindOne(ctx context.Context, id int64) (*TbRole, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tbRoleRows, m.table)
	var resp TbRole
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

func (m *defaultTbRoleModel) FindOneByCode(ctx context.Context, code string) (*TbRole, error) {
	var resp TbRole
	query := fmt.Sprintf("select %s from %s where `code` = ? limit 1", tbRoleRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, code)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTbRoleModel) Insert(ctx context.Context, data *TbRole) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, tbRoleRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Code, data.Name, data.Status, data.Remark, data.Permissions, data.CreateTs, data.UpdateTs)
	return ret, err
}

func (m *defaultTbRoleModel) Update(ctx context.Context, newData *TbRole) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tbRoleRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Code, newData.Name, newData.Status, newData.Remark, newData.Permissions, newData.CreateTs, newData.UpdateTs, newData.Id)
	return err
}

func (m *defaultTbRoleModel) tableName() string {
	return m.table
}
