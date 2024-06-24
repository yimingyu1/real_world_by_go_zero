package user

import (
	"context"

	"leanr_go_zero/restful/internal/svc"
	"leanr_go_zero/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 登出
func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() (resp *types.UserResp, err error) {
	// todo: add your logic here and delete this line

	return
}
