package profile

import (
	"context"

	"leanr_go_zero/restful/internal/svc"
	"leanr_go_zero/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查看文章
func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleLogic) GetArticle(req *types.GetArticleReq) (resp *types.ArticleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
