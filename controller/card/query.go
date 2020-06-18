package card

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/liuximu/flashcard/service"
	"github.com/liuximu/flashcard/shared"
)

func One(ctx echo.Context) shared.Rsp {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return shared.ErrRspBadParam
	}

	learner := GetUser(ctx)
	ctx1 := shared.EchoCtx2LogCtx(ctx)

	one, err := service.Card.GetByID(ctx1, int64(id), learner)
	if err != nil {
		return shared.ErrRspSysFail
	}
	if one == nil {
		return shared.ErrRspBadParam
	}

	return shared.NewSuccJSONRsp(one)
}

func List(ctx echo.Context) shared.Rsp {
	kid, err := strconv.Atoi(ctx.Param("kid"))
	if err != nil {
		return shared.ErrRspBadParam
	}

	learner := GetUser(ctx)
	ctx1 := shared.EchoCtx2LogCtx(ctx)

	list, err := service.Card.GetListByKID(ctx1, int64(kid), learner, 0)
	if err != nil {
		return shared.ErrRspSysFail
	}

	return shared.NewSuccJSONRsp(list)
}

func GetUser(ctx echo.Context) *service.Learner {
	return ctx.Get("Learner").(*service.Learner)
}
