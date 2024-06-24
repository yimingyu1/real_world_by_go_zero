package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ArticleCommentModel = (*customArticleCommentModel)(nil)

type (
	// ArticleCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleCommentModel.
	ArticleCommentModel interface {
		articleCommentModel
		withSession(session sqlx.Session) ArticleCommentModel
	}

	customArticleCommentModel struct {
		*defaultArticleCommentModel
	}
)

// NewArticleCommentModel returns a model for the database table.
func NewArticleCommentModel(conn sqlx.SqlConn) ArticleCommentModel {
	return &customArticleCommentModel{
		defaultArticleCommentModel: newArticleCommentModel(conn),
	}
}

func (m *customArticleCommentModel) withSession(session sqlx.Session) ArticleCommentModel {
	return NewArticleCommentModel(sqlx.NewSqlConnFromSession(session))
}
