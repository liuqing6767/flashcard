package knows

import (
	"unicode/utf8"

	"github.com/labstack/echo/v4"
	"github.com/liuximu/flashcard/service"
	"github.com/liuximu/flashcard/shared"
)

func Create(ctx echo.Context) shared.Rsp {
	learner := GetUser(ctx)
	param := &struct {
		Pid  int64  `json:"pid"`
		Name string `json:"name"`
	}{}
	if err := ctx.Bind(param); err != nil {
		return shared.ErrRspBadParam
	}

	if utf8.RuneCountInString(param.Name) < 1 || utf8.RuneCountInString(param.Name) > 10 {
		return shared.NewBadParam("知识点名称长度必须为1-10")
	}

	ctx1 := shared.EchoCtx2LogCtx(ctx)
	if pid := param.Pid; pid != 0 {
		parentNode, err := service.Know.GetOneByID(ctx1, learner, pid)
		if err != nil {
			return shared.ErrRspSysFail
		}
		if parentNode == nil {
			return shared.ErrRspBadParam

		}
	}

	node, err := service.Know.GetOneByName(ctx1, learner, param.Name)
	if err != nil {
		return shared.ErrRspSysFail
	}
	if node != nil {
		return shared.NewBadParam("知识点已存在")
	}

	node, err = service.Know.Create(ctx1, learner, param.Pid, param.Name)
	if err != nil {
		return shared.ErrRspSysFail
	}

	return shared.NewSuccJSONRsp(node)
}
