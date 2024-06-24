package profile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"leanr_go_zero/restful/internal/logic/profile"
	"leanr_go_zero/restful/internal/svc"
	"leanr_go_zero/restful/internal/types"
)

// 关注文章
func FavoriteArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteArticleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := profile.NewFavoriteArticleLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteArticle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
