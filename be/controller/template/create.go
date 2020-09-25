package template

import (
	"encoding/json"
	"unicode/utf8"

	"github.com/labstack/echo/v4"
	"github.com/liuximu/flashcard/be/service"
	"github.com/liuximu/flashcard/be/shared"
)

func Create(ctx echo.Context) shared.Rsp {
	param := &service.CardTemplate{}
	if err := ctx.Bind(param); err != nil {
		return shared.ErrRspBadParam
	}
	if err := validate(param); err != nil {
		return err
	}

	ctx1 := shared.EchoCtx2LogCtx(ctx)
	learner := GetUser(ctx)
	param.UID = learner.ID

	if err := param.Create(ctx1); err != nil {
		return shared.ErrRspSysFail
	}

	return shared.NewSuccJSONRsp(nil)
}

func validate(param *service.CardTemplate) shared.Rsp {
	if utf8.RuneCountInString(param.Name) < 1 || utf8.RuneCountInString(param.Name) > 10 {
		return shared.NewBadParam("模板名称长度必须为1-10")
	}

	tmp := map[string]interface{}{}
	if err := json.Unmarshal([]byte(param.DataDemo), &tmp); err != nil {
		return shared.ErrRspBadParam
	}

	data, ok := tmp[param.CardNameKey]
	if !ok {
		return shared.ErrRspBadParam
	}
	_, ok = data.(string)
	if !ok {
		return shared.ErrRspBadParam
	}

	return nil
}
