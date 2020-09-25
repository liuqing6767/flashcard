package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/liuximu/flashcard/be/service"
	"github.com/liuximu/flashcard/be/shared"
)

func MustLogin(ctx echo.Context) error {
	// ctx.Set("Learner", &service.Learner{
	// 	ID: 1,
	// })
	// return nil

	tokenCookie, err := ctx.Cookie("token")
	if err != nil {
		shared.ErrRspUserNoLogin.Render(ctx)
		return err
	}
	ctx1 := shared.EchoCtx2LogCtx(ctx)
	learner, err := service.Auth.GetUserByToken(ctx1, tokenCookie.Value)
	if err != nil {
		shared.ErrRspUserNoLogin.Render(ctx)
		return err
	}

	ctx.Set("Learner", learner)

	token, err := learner.GenToken(shared.TokenTypeLogin)
	if err != nil {
		shared.ErrRspUserNoLogin.Render(ctx)
		ctx.Logger().Error(err)
		return err
	}

	sendToken(ctx, token)

	return nil
}

func MustLoginHandlerFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if err := MustLogin(ctx); err != nil {
			ctx.Error(err)
			return err
		}

		return next(ctx)
	}
}
