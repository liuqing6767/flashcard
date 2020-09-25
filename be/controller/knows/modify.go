package knows

import (
	"unicode/utf8"

	"github.com/labstack/echo/v4"
	"github.com/liuximu/flashcard/be/dao"
	"github.com/liuximu/flashcard/be/service"
	"github.com/liuximu/flashcard/be/shared"
)

func Modify(ctx echo.Context) shared.Rsp {
	learner := GetUser(ctx)
	param := &struct {
		Type     string `json:"type"`
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		RID      int64  `json:"rid"`
		DropType string `json:"drop_type"`
	}{}
	if err := ctx.Bind(param); err != nil {
		return shared.ErrRspBadParam
	}

	ctx1 := shared.EchoCtx2LogCtx(ctx)

	switch param.Type {
	case "rename":
		if utf8.RuneCountInString(param.Name) < 1 || utf8.RuneCountInString(param.Name) > 10 {
			return shared.NewBadParam("知识点名称长度必须为1-10")
		}
		affecteds, err := service.Know.Update(ctx1, learner, param.ID, &dao.KnowPointParam{
			Name: &param.Name,
		})
		if err != nil {
			return shared.ErrRspSysFail
		}
		if affecteds == 0 {
			return shared.ErrRspBadParam
		}
	case "delete":
		affecteds, err := service.Know.Delete(ctx1, learner, param.ID)
		if err != nil {
			return shared.ErrRspSysFail
		}
		if affecteds == 0 {
			return shared.ErrRspBadParam
		}
	case "drag":
		allowedType := map[string]bool{
			"inner":  true,
			"before": true,
			"after":  true,
		}
		if !allowedType[param.DropType] {
			return shared.ErrRspBadParam
		}
		if err := service.Know.Drap(ctx1, learner, param.ID, param.RID, param.DropType); err != nil {
			return shared.ErrRspSysFail
		}
	default:
		return shared.ErrRspBadParam
	}
	return shared.NewSuccJSONRsp(nil)
}
