package profile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"leanr_go_zero/restful/internal/logic/profile"
	"leanr_go_zero/restful/internal/svc"
	"leanr_go_zero/restful/internal/types"
)

// 查看用户主页
func GetProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProfileReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := profile.NewGetProfileLogic(r.Context(), svcCtx)
		resp, err := l.GetProfile(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
