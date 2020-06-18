package auth

import (
	"regexp"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/liuximu/flashcard/dao"
	"github.com/liuximu/flashcard/service"
	"github.com/liuximu/flashcard/shared"
)

var emailReg = regexp.MustCompile(`^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`)

func Register(ctx echo.Context) shared.Rsp {
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
		if learner, err = service.Auth.CreateUser(ctx1, param.Email, param.Password); err != nil {
			return shared.ErrRspSysFail
		}
	}

	if learner.Status != dao.LearnerStatusRegisting {
		return shared.NewFailJSONRsp(400, "邮箱已经注册，请登录或者找回密码")
	}

	now := time.Now()
	if learner, err = service.Auth.UpdateUser(ctx1, param.Email, &dao.AuthLearnerParam{
		Password:     &param.Password,
		RegisterTime: &now,
		Status:       &dao.LearnerStatusRegisting,
	}); err != nil {
		return shared.ErrRspSysFail
	}

	if err := learner.SendRegisterEmail(ctx1); err != nil {
		return shared.ErrRspSysFail
	}

	return shared.NewSuccJSONRsp(nil, "注册成功，请点击邮件中的链接完成激活")
}

func RegisterActive(ctx echo.Context) shared.Rsp {
	token := ctx.QueryParam("token")
	ctx1 := shared.EchoCtx2LogCtx(ctx)
	learner, err := service.Auth.GetUserByToken(ctx1, token)
	if err != nil {
		return shared.ErrRspSysFail
	}

	if time.Now().Sub(learner.Token.CreateTime) > time.Hour*24 {
		return shared.NewFailJSONRsp(400, "链接已过期，请重新注册")
	}

	if learner == nil {
		return shared.ErrRspBadParam
	}

	if learner, err = service.Auth.UpdateUser(ctx1, learner.Email, &dao.AuthLearnerParam{
		Status: &dao.LearnerStatuOk,
	}); err != nil {
		return shared.ErrRspSysFail
	}

	// tokenLogin, err := learner.Token(service.TokenTypeLogin)
	// if err != nil {
	// 	return shared.ErrRspSysFail
	// }
	// sendToken(ctx, tokenLogin)
	toHomePage(ctx)

	return nil
}

func toHomePage(ctx echo.Context) {
	ctx.Redirect(302, shared.AppConfSingleton.Links.LoginPage)
}
