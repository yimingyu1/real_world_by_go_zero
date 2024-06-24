package comment

import (
	"context"

	"leanr_go_zero/restful/internal/svc"
	"leanr_go_zero/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取tagList
func NewGetTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagListLogic {
	return &GetTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTagListLogic) GetTagList() (resp *types.TagListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
