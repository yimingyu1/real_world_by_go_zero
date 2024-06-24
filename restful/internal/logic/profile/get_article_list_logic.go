package profile

import (
	"context"

	"leanr_go_zero/restful/internal/svc"
	"leanr_go_zero/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取文章列表
func NewGetArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleListLogic {
	return &GetArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleListLogic) GetArticleList(req *types.GetArticleListReq) (resp *types.ArticleListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
