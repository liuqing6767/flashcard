package template

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/liuximu/flashcard/dao"
	"github.com/liuximu/flashcard/service"
	"github.com/liuximu/flashcard/shared"
)

func Delete(ctx echo.Context) shared.Rsp {
	tid, err := strconv.Atoi(ctx.Param("tid"))
	if err != nil {
		return shared.ErrRspBadParam
	}

	ctx1 := shared.EchoCtx2LogCtx(ctx)
	temp, err := service.Temp.GetOneByID(ctx1, int64(tid))
	if err != nil {
		return shared.ErrRspSysFail
	}
	learner := GetUser(ctx)
	if temp == nil || temp.UID != learner.ID {
		return shared.ErrRspBadParam
	}

	temp.Status = dao.DeletedYes

	if err := temp.Update(ctx1); err != nil {
		return shared.ErrRspSysFail
	}

	return shared.NewSuccJSONRsp(nil)
}
