package auth

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/liuximu/flashcard/dao"
	"github.com/liuximu/flashcard/service"
	"github.com/liuximu/flashcard/shared"
)

type LoginParam struct {
	Email    string
	Password string
}

func Login(ctx echo.Context) shared.Rsp {
	param := &LoginParam{}
	if err := ctx.Bind(param); err != nil {
		return shared.ErrRspBadParam
	}
	if !emailReg.MatchString(param.Email) {
		return shared.NewFailJSONRsp(422, "邮箱格式错误")
	}

	ctx1 := shared.EchoCtx2LogCtx(ctx)
	learner, err := service.Auth.GetUserByEmail(ctx1, param.Email)
	if err != nil {
		return shared.ErrRspSysFail
	}
	if learner == nil {
		return shared.NewFailJSONRsp(422, "用户名未注册，请先注册")
	}

	if !learner.RightPassword(param.Password) {
		return shared.NewFailJSONRsp(422, "用户名或者密码错误")
	}

	if learner.Status == dao.LearnerStatusRegisting {
		if err := learner.SendRegisterEmail(ctx1); err != nil {
			return shared.ErrRspSysFail
		}

		return shared.NewFailJSONRsp(422, "用户名未激活，请打开邮箱点击激活链接")
	}

	token, err := learner.GenToken(shared.TokenTypeLogin)
	if err != nil {
		return shared.ErrRspSysFail
	}
	sendToken(ctx, token)

	return shared.NewSuccJSONRsp(map[string]string{
		"email": learner.Email,
	})
}

func sendToken(ctx echo.Context, token string) {
	// 设置TOKEN
	ctx.SetCookie(&http.Cookie{
		Name:   "token",
		Value:  token,
		MaxAge: 3600 * 24 * 15,
	})

}
