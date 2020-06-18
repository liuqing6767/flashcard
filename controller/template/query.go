package template

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/liuximu/flashcard/service"
	"github.com/liuximu/flashcard/shared"
)

func List(ctx echo.Context) shared.Rsp {
	learner := GetUser(ctx)
	ctx1 := shared.EchoCtx2LogCtx(ctx)
	list, err := service.Temp.GetListByLearner(ctx1, learner)
	if err != nil {
		return shared.ErrRspSysFail
	}

	return shared.NewSuccJSONRsp(list)
}

func One(ctx echo.Context) shared.Rsp {
	tid, err := strconv.Atoi(ctx.Param("tid"))
	if err != nil {
		return shared.ErrRspBadParam
	}
	learner := GetUser(ctx)

	ctx1 := shared.EchoCtx2LogCtx(ctx)
	temp, err := service.Temp.GetOneByID(ctx1, int64(tid))
	if err != nil {
		return shared.ErrRspSysFail
	}
	if temp == nil {
		return shared.ErrRspBadParam
	}

	if !temp.IsOfficial && temp.UID != learner.ID {
		return shared.ErrRspBadParam
	}

	return shared.NewSuccJSONRsp(temp)
}

func GetUser(ctx echo.Context) *service.Learner {
	return ctx.Get("Learner").(*service.Learner)
}
