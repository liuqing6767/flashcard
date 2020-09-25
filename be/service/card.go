package service

import (
	"time"

	"github.com/liuximu/flashcard/be/dao"
	"github.com/liuximu/flashcard/be/shared"
)

type card struct{}

var Card = &card{}

type CardDetail struct {
	ID         int64  `json:"id"`
	UID        int64  `json:"uid"`
	Name       string `json:"name"`
	Data       string `json:"data"`
	KnowID     int64  `json:"know_id"`
	TemplateID int64  `json:"template_id"`

	NextLearnTime int64 `json:"next_learn_time"`
	Progress      int8  `json:"progress"`
	ProgressRate  int8  `json:"progress_rate"`
}

func (cd *CardDetail) GetMemeoryCurve(ctx shared.Context) (*MemoryCurve, error) {
	return Memory.GetByID(ctx, 0)
}

func (cd *CardDetail) toCardParam() *dao.CardParam {
	param := &dao.CardParam{
		Uid:        &cd.UID,
		Name:       &cd.Name,
		Data:       &cd.Data,
		KnowID:     &cd.KnowID,
		TemplateID: &cd.TemplateID,
	}
	if cd.ID != 0 {
		param.Id = &cd.ID
	}

	return param
}

func (cd *CardDetail) Create(ctx shared.Context) error {
	memoryCurve, err := Memory.GetByID(ctx, 0)
	if err != nil {
		ctx.Error(err)
		return err
	}

	_, nextLearnTime := memoryCurve.Next(0, 3)
	param := cd.toCardParam()
	if !nextLearnTime.IsZero() {
		param.NextLearnTime = &nextLearnTime
	}

	rst, err := param.Create(ctx, dao.XimuFlashcard())
	if err != nil {
		ctx.Error(err)
		return err
	}
	cd.ID, err = rst.LastInsertId()
	if err != nil {
		ctx.Error(err)
		return err
	}

	return nil
}

func (cd *CardDetail) UpdateProgress(ctx shared.Context) error {
	t := time.Unix(cd.NextLearnTime, 0)
	param := &dao.CardParam{
		Progress:      &cd.Progress,
		NextLearnTime: &t,
	}
	_, err := param.Update(ctx, dao.XimuFlashcard(), &dao.CardParam{
		Id: &cd.ID,
	})
	if err != nil {
		ctx.Error(err)
		return err
	}

	return nil
}

func (cd *CardDetail) Update(ctx shared.Context) error {
	param := cd.toCardParam()
	_, err := param.Update(ctx, dao.XimuFlashcard(), &dao.CardParam{
		Id: &cd.ID,
	})
	if err != nil {
		ctx.Error(err)
		return err
	}

	return nil
}

func (cd *CardDetail) Delete(ctx shared.Context, learner *Learner) (int64, error) {
	param := &dao.CardParam{
		Id:  &cd.ID,
		Uid: &learner.ID,
	}
	rst, err := param.Delete(ctx, dao.XimuFlashcard())
	if err != nil {
		ctx.Error(err)
		return 0, err
	}
	row, err := rst.RowsAffected()
	if err != nil {
		ctx.Error(err)
		return 0, err
	}
	return row, nil
}

func (c *card) GetLeariningCount(ctx shared.Context, learner *Learner) (int64, error) {
	count, err := dao.CardCountLearning(ctx, dao.XimuFlashcard(), learner.ID, time.Now())
	if err != nil {
		ctx.Error(err)
		return 0, err
	}
	return count, nil
}

func (c *card) GetByID(ctx shared.Context, id int64, learner *Learner) (*CardDetail, error) {
	one := &dao.Card{}
	err := one.Query(ctx, dao.XimuFlashcard(), &dao.CardParam{
		Id:  &id,
		Uid: &learner.ID,
	}, nil)
	if err != nil {
		if shared.EmptyResult(err) {
			return nil, nil
		}
		ctx.Error(err)
		return nil, err
	}
	return newCardDetail(one), nil
}

func (c *card) GetListByKID(ctx shared.Context, where *dao.CardParam) ([]*CardDetail, error) {
	list := dao.CardList{}
	order := "next_learn_time asc"
	where.OrderBy = &order

	if err := list.Query(ctx, dao.XimuFlashcard(), where, nil); err != nil {
		ctx.Error(err)
		return nil, err
	}

	ids := []int64{}
	for _, one := range list {
		ids = append(ids, one.Id)
	}

	rst := []*CardDetail{}
	for _, one := range list {
		rst = append(rst, newCardDetail(one))
	}

	return rst, nil
}

func newCardDetail(one *dao.Card) *CardDetail {
	return &CardDetail{
		ID:         one.Id,
		UID:        one.Uid,
		Name:       one.Name,
		Data:       one.Data,
		KnowID:     one.KnowID,
		TemplateID: one.TemplateID,

		NextLearnTime: one.NextLearnTime.Unix(),
		Progress:      one.Progress,
		ProgressRate:  int8(int64(one.Progress) * 100 / 7),
	}
}
