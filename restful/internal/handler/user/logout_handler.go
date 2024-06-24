package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"leanr_go_zero/restful/internal/logic/user"
	"leanr_go_zero/restful/internal/svc"
)

// 登出
func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
