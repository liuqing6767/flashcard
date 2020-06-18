package card

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"github.com/liuximu/flashcard/service"
	"github.com/liuximu/flashcard/shared"
)

func Upsert(ctx echo.Context) shared.Rsp {
	param := &service.CardDetail{}
	if err := ctx.Bind(param); err != nil {
		return shared.ErrRspBadParam
	}

	learner := GetUser(ctx)
	ctx1 := shared.EchoCtx2LogCtx(ctx)

	param.UID = learner.ID

	template, err := service.Temp.GetOneByID(ctx1, param.TemplateID)
	if err != nil {
		return shared.ErrRspSysFail
	}
	if template == nil {
		return shared.ErrRspBadParam
	}

	if param.Name == "" {
		tmp := map[string]interface{}{}
		if err := json.Unmarshal([]byte(param.Data), &tmp); err != nil {
			return shared.ErrRspBadParam
		}

		tmpName, ok := tmp[template.CardNameKey]
		if !ok {
			return shared.ErrRspBadParam
		}
		param.Name, ok = tmpName.(string)
		if !ok {
			return shared.ErrRspBadParam
		}
	}

	know, err := service.Know.GetOneByID(ctx1, learner, param.KnowID)
	if err != nil {
		return shared.ErrRspSysFail
	}
	if know == nil {
		return shared.ErrRspBadParam
	}

	if param.ID != 0 {
		err = param.Update(ctx1)
	} else {
		err = param.Create(ctx1)
	}
	if err != nil {
		return shared.ErrRspSysFail
	}

	return shared.NewSuccJSONRsp(nil)
}
