package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ FollowModel = (*customFollowModel)(nil)

type (
	// FollowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowModel.
	FollowModel interface {
		followModel
		withSession(session sqlx.Session) FollowModel
	}

	customFollowModel struct {
		*defaultFollowModel
	}
)

// NewFollowModel returns a model for the database table.
func NewFollowModel(conn sqlx.SqlConn) FollowModel {
	return &customFollowModel{
		defaultFollowModel: newFollowModel(conn),
	}
}

func (m *customFollowModel) withSession(session sqlx.Session) FollowModel {
	return NewFollowModel(sqlx.NewSqlConnFromSession(session))
}
