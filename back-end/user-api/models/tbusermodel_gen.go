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
	tbUserFieldNames          = builder.RawFieldNames(&TbUser{})
	tbUserRows                = strings.Join(tbUserFieldNames, ",")
	tbUserRowsExpectAutoSet   = strings.Join(stringx.Remove(tbUserFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tbUserRowsWithPlaceHolder = strings.Join(stringx.Remove(tbUserFieldNames, "`user_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tbUserModel interface {
		Insert(ctx context.Context, data *TbUser) (sql.Result, error)
		FindOne(ctx context.Context, userId int64) (*TbUser, error)
		FindOneByUsername(ctx context.Context, username string) (*TbUser, error)
		Update(ctx context.Context, data *TbUser) error
		Delete(ctx context.Context, userId int64) error
	}

	defaultTbUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TbUser struct {
		UserId   int64  `db:"user_id"`
		Username string `db:"username"`
		Password string `db:"password"`
		Nickname string `db:"nickname"`
		Avatar   string `db:"avatar"`
		Roles    string `db:"roles"`
		CreateTs int64  `db:"create_ts"`
		UpdateTs int64  `db:"update_ts"`
	}
)

func newTbUserModel(conn sqlx.SqlConn) *defaultTbUserModel {
	return &defaultTbUserModel{
		conn:  conn,
		table: "`tb_user`",
	}
}

func (m *defaultTbUserModel) Delete(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("delete from %s where `user_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, userId)
	return err
}

func (m *defaultTbUserModel) FindOne(ctx context.Context, userId int64) (*TbUser, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", tbUserRows, m.table)
	var resp TbUser
	err := m.conn.QueryRowCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTbUserModel) FindOneByUsername(ctx context.Context, username string) (*TbUser, error) {
	var resp TbUser
	query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", tbUserRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, username)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTbUserModel) Insert(ctx context.Context, data *TbUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, tbUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Username, data.Password, data.Nickname, data.Avatar, data.Roles, data.CreateTs, data.UpdateTs)
	return ret, err
}

func (m *defaultTbUserModel) Update(ctx context.Context, newData *TbUser) error {
	query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, tbUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Username, newData.Password, newData.Nickname, newData.Avatar, newData.Roles, newData.CreateTs, newData.UpdateTs, newData.UserId)
	return err
}

func (m *defaultTbUserModel) tableName() string {
	return m.table
}
