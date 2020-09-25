package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/liuximu/flashcard/be/shared"
)

// TODO
func ResetEmail(ctx echo.Context) shared.Rsp {
	return shared.NewSuccJSONRsp(nil)
	return shared.NewFailJSONRsp(400, "账号或密码错误")
}
