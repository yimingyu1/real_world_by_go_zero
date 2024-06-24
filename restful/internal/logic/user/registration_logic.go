package user

import (
	"context"
	"leanr_go_zero/common"
	"leanr_go_zero/common/model"
	"time"

	"leanr_go_zero/restful/internal/svc"
	"leanr_go_zero/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegistrationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 注册用户
func NewRegistrationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegistrationLogic {
	return &RegistrationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegistrationLogic) Registration(req *types.RegistrationReq) (resp *types.UserResp, err error) {
	// 查询用户是否存在
	email, _ := l.svcCtx.UserModel.FindOneByEmail(l.ctx, req.Email)
	if email != nil {
		return &types.UserResp{
			BaseResp: types.BaseResp{
				Code: common.EMAIL_REGISTERED_ERROR.Code,
				Msg:  common.EMAIL_REGISTERED_ERROR.Message,
			},
			Data: types.User{},
		}, nil
	}
	// 如果用户不存在，创建用户
	now := time.Now()
	result, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		UserName:   req.UserName,
		Email:      req.Email,
		Password:   req.Password,
		Bio:        req.Bio,
		DelState:   model.DEL_STATE_NO,
		CreateTime: &now,
		UpdateTime: &now,
	})
	if err != nil {
		return &types.UserResp{
			BaseResp: types.BaseResp{
				Code: common.INSERT_USER_ERROR.Code,
				Msg:  common.INSERT_USER_ERROR.Message,
			},
			Data: types.User{},
		}, err
	}
	//
	return
}
