package knows

import (
	"github.com/labstack/echo/v4"
	"github.com/liuximu/flashcard/be/service"
	"github.com/liuximu/flashcard/be/shared"
)

func Tree(ctx echo.Context) shared.Rsp {
	learner := GetUser(ctx)
	ctx1 := shared.EchoCtx2LogCtx(ctx)
	root, err := service.Know.GetKnowTree(ctx1, learner, 0)
	if err != nil {
		return shared.ErrRspSysFail
	}

	root.ValueCount(ctx1, learner)

	return shared.NewSuccJSONRsp(root)
}

func GetUser(ctx echo.Context) *service.Learner {
	return ctx.Get("Learner").(*service.Learner)
}
