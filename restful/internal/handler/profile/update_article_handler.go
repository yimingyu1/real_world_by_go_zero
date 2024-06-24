package profile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"leanr_go_zero/restful/internal/logic/profile"
	"leanr_go_zero/restful/internal/svc"
	"leanr_go_zero/restful/internal/types"
)

// 更新文章
func UpdateArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateArticleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := profile.NewUpdateArticleLogic(r.Context(), svcCtx)
		resp, err := l.UpdateArticle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
