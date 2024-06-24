package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"leanr_go_zero/common/core"
	"leanr_go_zero/common/model"
	"leanr_go_zero/restful/internal/config"
	"leanr_go_zero/restful/internal/middleware"
)

type ServiceContext struct {
	Config                   config.Config
	AuthenticationMiddleware rest.Middleware
	Redis                    *redis.Redis
	UserModel                model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                   c,
		AuthenticationMiddleware: middleware.NewAuthenticationMiddleware().Handle,
		Redis:                    core.NewRedis(c.Redis.Host, c.Redis.Type),
		UserModel:                model.NewUserModel(core.NewSqlConn(c.DB.DataSource)),
	}
}
