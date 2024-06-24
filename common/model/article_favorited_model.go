package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ArticleFavoritedModel = (*customArticleFavoritedModel)(nil)

type (
	// ArticleFavoritedModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleFavoritedModel.
	ArticleFavoritedModel interface {
		articleFavoritedModel
		withSession(session sqlx.Session) ArticleFavoritedModel
	}

	customArticleFavoritedModel struct {
		*defaultArticleFavoritedModel
	}
)

// NewArticleFavoritedModel returns a model for the database table.
func NewArticleFavoritedModel(conn sqlx.SqlConn) ArticleFavoritedModel {
	return &customArticleFavoritedModel{
		defaultArticleFavoritedModel: newArticleFavoritedModel(conn),
	}
}

func (m *customArticleFavoritedModel) withSession(session sqlx.Session) ArticleFavoritedModel {
	return NewArticleFavoritedModel(sqlx.NewSqlConnFromSession(session))
}
