package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

const DEL_STATE_NO int64 = 0  //未删除
const DEL_STATE_YES int64 = 1 //已删除
