package profile

import (
	"context"

	"leanr_go_zero/restful/internal/svc"
	"leanr_go_zero/restful/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新文章
func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateArticleLogic) UpdateArticle(req *types.UpdateArticleReq) (resp *types.ArticleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
