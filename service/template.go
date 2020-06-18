package service

import (
	"github.com/liuximu/flashcard/dao"
	"github.com/liuximu/flashcard/shared"
)

type temp struct{}

var Temp = &temp{}

type CardTemplate struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	ContentA    string `json:"content_a"`
	ContentB    string `json:"content_b"`
	DataDemo    string `json:"data_demo"`
	CardNameKey string `json:"card_name_key"`
	UID         int64  `json:"uid"`
	IsOfficial  bool   `json:"is_official"`
	CreateAt    string `json:"create_at"`
	Status      int8   `json:"-"`
}

func (ct *CardTemplate) Update(ctx shared.Context) error {
	template := &dao.CardTemplateParam{
		Name:        &ct.Name,
		ContentA:    &ct.ContentA,
		ContentB:    &ct.ContentB,
		DataDemo:    &ct.DataDemo,
		CardNameKey: &ct.CardNameKey,
		Uid:         &ct.UID,
		Status:      &ct.Status,
	}

	rst, err := template.Update(ctx, dao.XimuFlashcard(), &dao.CardTemplateParam{
		Id: &ct.ID,
	})
	if err != nil {
		ctx.Error(err)
		return err
	}

	id, err := rst.LastInsertId()
	if err != nil {
		ctx.Error(err)
		return err
	}

	ct.ID = id
	return nil
}

func (ct *CardTemplate) Create(ctx shared.Context) error {
	template := &dao.CardTemplateParam{
		Name:        &ct.Name,
		ContentA:    &ct.ContentA,
		ContentB:    &ct.ContentB,
		DataDemo:    &ct.DataDemo,
		CardNameKey: &ct.CardNameKey,
		Uid:         &ct.UID,
	}

	rst, err := template.Create(ctx, dao.XimuFlashcard())
	if err != nil {
		ctx.Error(err)
		return err
	}

	id, err := rst.LastInsertId()
	if err != nil {
		ctx.Error(err)
		return err
	}

	ct.ID = id
	return nil
}

func (t *temp) GetOneByID(ctx shared.Context, id int64) (*CardTemplate, error) {
	return t.getOne(ctx, &dao.CardTemplateParam{
		Id:     &id,
		Status: &dao.DeletedNo,
	})
}

func (t *temp) getOne(ctx shared.Context, where *dao.CardTemplateParam) (*CardTemplate, error) {
	one := &dao.CardTemplate{}
	if err := one.Query(ctx, dao.XimuFlashcard(), where, nil); err != nil {
		if shared.EmptyResult(err) {
			return nil, nil
		}
		ctx.Error(err)
		return nil, err
	}

	return newCardTemplate(one), nil
}

func newCardTemplate(one *dao.CardTemplate) *CardTemplate {
	return &CardTemplate{
		ID:          one.Id,
		Name:        one.Name,
		ContentA:    one.ContentA,
		ContentB:    one.ContentB,
		DataDemo:    one.DataDemo,
		CardNameKey: one.CardNameKey,
		UID:         one.Uid,
		IsOfficial:  one.Uid == 0,
	}
}

func (t *temp) GetListByLearner(ctx shared.Context, learner *Learner) ([]*CardTemplate, error) {
	list := dao.CardTemplateList{}
	err := list.Query(ctx, dao.XimuFlashcard(), &dao.CardTemplateParam{
		Status: &dao.DeletedNo,
		Uids:   []int64{0, learner.ID},
	}, nil)
	if err != nil {
		ctx.Error(err)
		return nil, err
	}

	rst := []*CardTemplate{}
	for _, one := range list {
		rst = append(rst, newCardTemplate(one))
	}
	return rst, err
}

func (t *temp) GetListByIDs(ctx shared.Context, ids []int64) ([]*CardTemplate, error) {
	list := dao.CardTemplateList{}
	err := list.Query(ctx, dao.XimuFlashcard(), &dao.CardTemplateParam{
		Ids: ids,
	}, nil)
	if err != nil {
		ctx.Error(err)
		return nil, err
	}

	rst := []*CardTemplate{}
	for _, one := range list {
		rst = append(rst, newCardTemplate(one))
	}
	return rst, err
}
