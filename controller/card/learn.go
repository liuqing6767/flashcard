package card

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/liuximu/flashcard/dao"
	"github.com/liuximu/flashcard/service"
	"github.com/liuximu/flashcard/shared"
)

func LearnProgressStatis(ctx echo.Context) shared.Rsp {
	learner := GetUser(ctx)
	ctx1 := shared.EchoCtx2LogCtx(ctx)
	count, err := service.Card.GetLeariningCount(ctx1, learner)
	if err != nil {
		return shared.ErrRspSysFail
	}

	return shared.NewSuccJSONRsp(map[string]interface{}{
		"count": count,
	})
}

func UpdateLearnProgress(ctx echo.Context) shared.Rsp {
	param := &struct {
		CardID      int64 `json:"cid"`
		Familiarity int8  `json:"f"`
	}{}
	if err := ctx.Bind(param); err != nil {
		return shared.ErrRspBadParam
	}

	if param.Familiarity < 0 || param.Familiarity > 3 {
		return shared.ErrRspBadParam
	}

	learner := GetUser(ctx)
	ctx1 := shared.EchoCtx2LogCtx(ctx)
	card, err := service.Card.GetByID(ctx1, param.CardID, learner)
	if err != nil {
		return shared.ErrRspSysFail
	}
	if card == nil {
		return shared.ErrRspBadParam
	}

	mc, err := card.GetMemeoryCurve(ctx1)
	if err != nil {
		return shared.ErrRspSysFail
	}

	var nextLearnTime time.Time
	card.Progress, nextLearnTime = mc.Next(card.Progress, param.Familiarity)
	card.NextLearnTime = nextLearnTime.Unix()

	card.UpdateProgress(ctx1)

	return shared.NewSuccJSONRsp(nil)
}

func LearningList(ctx echo.Context) shared.Rsp {
	kid, err := strconv.Atoi(ctx.Param("kid"))
	if err != nil {
		return shared.ErrRspBadParam
	}

	cid, _ := strconv.Atoi(ctx.QueryParam("cid"))

	learner := GetUser(ctx)
	ctx1 := shared.EchoCtx2LogCtx(ctx)

	_kid := int64(kid)
	now := time.Now()
	where := &dao.CardParam{
		KnowID:           &_kid,
		Uid:              &learner.ID,
		MaxNextLearnTime: &now,
	}
	if cid != 0 {
		_cid := int64(cid)
		where.Id = &_cid
	}
	cardList, err := service.Card.GetListByKID(ctx1, where)
	if err != nil {
		return shared.ErrRspSysFail
	}

	if len(cardList) == 0 {
		return shared.NewSuccJSONRsp(map[string]interface{}{
			"cards":     cardList,
			"templates": nil,
		})
	}

	tidMap, tids := map[int64]bool{}, []int64{}
	for _, one := range cardList {
		if !tidMap[one.TemplateID] {
			tidMap[one.TemplateID] = true
			tids = append(tids, one.TemplateID)
		}
	}

	templateList, err := service.Temp.GetListByIDs(ctx1, tids)
	if err != nil {
		return shared.ErrRspSysFail
	}

	return shared.NewSuccJSONRsp(map[string]interface{}{
		"cards":     cardList,
		"templates": templateList,
	})
}
