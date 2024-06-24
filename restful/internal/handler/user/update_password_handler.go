package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"leanr_go_zero/restful/internal/logic/user"
	"leanr_go_zero/restful/internal/svc"
	"leanr_go_zero/restful/internal/types"
)

// 修改密码
func UpdatePasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePasswordReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewUpdatePasswordLogic(r.Context(), svcCtx)
		resp, err := l.UpdatePassword(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
