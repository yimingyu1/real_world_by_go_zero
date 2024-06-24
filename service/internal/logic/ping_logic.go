package logic

import (
	"context"

	"leanr_go_zero/service/internal/svc"
	"leanr_go_zero/service/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *service.Request) (*service.Response, error) {
	// todo: add your logic here and delete this line

	return &service.Response{}, nil
}
