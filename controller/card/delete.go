package card

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/liuximu/flashcard/service"
	"github.com/liuximu/flashcard/shared"
)

func Delete(ctx echo.Context) shared.Rsp {
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

	rows, err := one.Delete(ctx1, learner)
	if err != nil {
		return shared.ErrRspSysFail
	}
	if rows == 0 {
		return shared.ErrRspBadParam
	}

	return shared.NewSuccJSONRsp(nil)
}
