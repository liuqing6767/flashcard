package card

import (
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/liuximu/flashcard/service"
	"github.com/liuximu/flashcard/shared"
)

func WordQuery(ctx echo.Context) shared.Rsp {
	word := strings.TrimSpace(ctx.Param("word"))
	if word == "" || strings.Index(word, " ") != -1 {
		return shared.ErrRspBadParam
	}
	word = strings.ToLower(word)

	GetUser(ctx)
	ctx1 := shared.EchoCtx2LogCtx(ctx)
	one, err := service.Word.QueryOrCreate(ctx1, word)
	if err != nil {
		return shared.ErrRspSysFail
	}
	if one == nil {
		return shared.ErrRspBadParam
	}

	return shared.NewSuccJSONRsp(one)
}
