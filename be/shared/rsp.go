package shared

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MineHandlerFunc func(echo.Context) Rsp

type Rsp interface {
	Render(echo.Context)
}

type JSONRsp struct {
	Errno int64       `json:"errno"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

func (jr *JSONRsp) Render(ctx echo.Context) {
	ctx.JSON(http.StatusOK, jr)
	// ctx.Response().Committed = true
}

func NewSuccJSONRsp(data interface{}, msgOption ...string) Rsp {
	msg := ""
	if len(msgOption) == 1 {
		msg = msgOption[0]
	}
	return &JSONRsp{
		Data: data,
		Msg:  msg,
	}
}

func NewFailJSONRsp(no int64, msg string) Rsp {
	return &JSONRsp{
		Msg:   msg,
		Errno: no,
	}
}

var _ Rsp = &JSONRsp{}

func Mine2Handler(m MineHandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if rsp := m(ctx); rsp != nil {
			// auth.SetToken(ctx)/
			rsp.Render(ctx)
		}

		return nil
	}
}

func NewBadParam(info string) Rsp {
	return NewFailJSONRsp(422, info)
}

var (
	ErrRspSysFail     = NewFailJSONRsp(500, "系统异常，请稍后再试")
	ErrRspBadParam    = NewFailJSONRsp(422, "请求参数非法")
	ErrRspUserNoLogin = NewFailJSONRsp(401, "用户未登录")
)
