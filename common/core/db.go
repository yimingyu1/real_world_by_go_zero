package core

import "github.com/zeromicro/go-zero/core/stores/sqlx"

func NewSqlConn(datasource string) sqlx.SqlConn {
	return sqlx.NewMysql(datasource)
}
