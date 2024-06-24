// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	articleCommentFieldNames          = builder.RawFieldNames(&ArticleComment{})
	articleCommentRows                = strings.Join(articleCommentFieldNames, ",")
	articleCommentRowsExpectAutoSet   = strings.Join(stringx.Remove(articleCommentFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	articleCommentRowsWithPlaceHolder = strings.Join(stringx.Remove(articleCommentFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	articleCommentModel interface {
		Insert(ctx context.Context, data *ArticleComment) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ArticleComment, error)
		Update(ctx context.Context, data *ArticleComment) error
		Delete(ctx context.Context, id int64) error
	}

	defaultArticleCommentModel struct {
		conn  sqlx.SqlConn
		table string
	}

	ArticleComment struct {
		Id         int64          `db:"id"`          // id
		AutherId   int64          `db:"auther_id"`   // 评论作者
		ArticleId  int64          `db:"article_id"`  // 文章id
		Body       sql.NullString `db:"body"`        // 评论内容
		DelState   int64          `db:"del_state"`   // 删除状态 0：未删除 1：已删除
		CreateTime time.Time      `db:"create_time"` // 创建时间
		UpdateTime time.Time      `db:"update_time"` // 更新时间
		DeleteTime time.Time      `db:"delete_time"` // 删除时间
	}
)

func newArticleCommentModel(conn sqlx.SqlConn) *defaultArticleCommentModel {
	return &defaultArticleCommentModel{
		conn:  conn,
		table: "`article_comment`",
	}
}

func (m *defaultArticleCommentModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultArticleCommentModel) FindOne(ctx context.Context, id int64) (*ArticleComment, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", articleCommentRows, m.table)
	var resp ArticleComment
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

func (m *defaultArticleCommentModel) Insert(ctx context.Context, data *ArticleComment) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, articleCommentRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.AutherId, data.ArticleId, data.Body, data.DelState, data.DeleteTime)
	return ret, err
}

func (m *defaultArticleCommentModel) Update(ctx context.Context, data *ArticleComment) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, articleCommentRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.AutherId, data.ArticleId, data.Body, data.DelState, data.DeleteTime, data.Id)
	return err
}

func (m *defaultArticleCommentModel) tableName() string {
	return m.table
}
