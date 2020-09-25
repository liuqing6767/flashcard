package card

import (
	"encoding/json"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/liuximu/flashcard/be/dao"
	"github.com/liuximu/flashcard/be/service"
	"github.com/liuximu/flashcard/be/shared"
)

const nodeName = "智能单词本"

// 谷歌浏览器扩展使用的接口
func UpsertExt(ctx echo.Context) shared.Rsp {
	word := strings.TrimSpace(ctx.QueryParam("word"))
	if word == "" || strings.Index(word, " ") != -1 {
		return shared.ErrRspBadParam
	}
	word = strings.ToLower(word)
	ctx1 := shared.EchoCtx2LogCtx(ctx)
	one, err := service.Word.QueryOrCreate(ctx1, word)
	if err != nil {
		return shared.ErrRspSysFail
	}
	if one == nil {
		return shared.ErrRspBadParam
	}

	learner := GetUser(ctx)

	// 看看用户是否有这个知识点
	knowNode, err := service.Know.GetOneByName(ctx1, learner, nodeName)
	if err != nil {
		return shared.ErrRspSysFail
	}
	// 创建
	if knowNode == nil {
		knowNode, err = service.Know.Create(ctx1, learner, 0, nodeName)
		if err != nil {
			return shared.ErrRspSysFail
		}
	}

	// 看看这个用户是否自己有了这个单词
	card, err := service.Card.GetByName(ctx1, word, learner)
	if err != nil {
		return shared.ErrRspSysFail
	}
	if card != nil {
		// 存在了，直接退出
		return shared.NewSuccJSONRsp(nil)
	}

	wordBs, err := json.Marshal(map[string]interface{}{
		"Word":   one.Word,
		"Detail": one,
	})
	if err != nil {
		return shared.ErrRspBadParam
	}

	param := &service.CardDetail{
		UID:        learner.ID,
		Name:       one.Word,
		Data:       string(wordBs),
		KnowID:     knowNode.ID,
		TemplateID: dao.WordTemplateID,
	}

	if err = param.Create(ctx1); err != nil {
		return shared.ErrRspSysFail
	}

	return shared.NewSuccJSONRsp(nil)
}
