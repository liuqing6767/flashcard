package service

import (
	"time"

	"github.com/liuximu/flashcard/dao"
	"github.com/liuximu/flashcard/shared"
)

type know struct{}

var Know = &know{}

type KnowNode struct {
	ID            int64       `json:"id"`
	PID           int64       `json:"pid"`
	Label         string      `json:"label"`
	TotalCount    int         `json:"total"`
	LearningCount int         `json:"learning"`
	Index         int64       `json:"index"`
	Children      []*KnowNode `json:"children"`
}

func (k *know) GetKnowTree(ctx shared.Context, learner *Learner, parentID int64) (*KnowNode, error) {
	knowList := dao.KnowPointList{}
	order := "inde asc"
	where := &dao.KnowPointParam{
		Uid:     &learner.ID,
		Status:  &dao.DeletedNo,
		OrderBy: &order,
	}
	if parentID != 0 {
		where.Pid = &parentID
	}

	if err := knowList.Query(ctx, dao.XimuFlashcard(), where, nil); err != nil {
		ctx.Error(err)
		return nil, err
	}

	root := &KnowNode{
		Label: "我的知识体系",
	}

	nodeMap := map[int64]*KnowNode{
		0: root,
	}

	for _, know := range knowList {
		node := nodeMap[know.Id]
		if node == nil {
			node = &KnowNode{}

			nodeMap[know.Id] = node
		}
		node.Label = know.Name
		node.ID = know.Id
		node.Index = know.Index
		node.PID = know.Pid

		parent := nodeMap[know.Pid]
		if parent == nil {
			parent = &KnowNode{}
			nodeMap[know.Pid] = parent
		}
		parent.Children = append(parent.Children, node)
	}

	return nodeMap[parentID], nil
}

func (k *KnowNode) ValueCount(ctx shared.Context, learner *Learner) error {
	cardList := dao.CardList{}
	if err := cardList.Query(ctx, dao.XimuFlashcard(), &dao.CardParam{
		Uid: &learner.ID,
	}, []string{"kid, next_learn_time"}); err != nil {
		ctx.Error(err)
		return err
	}
	all, learning := map[int64]int{}, map[int64]int{}
	now := time.Now()
	for _, one := range cardList {
		all[one.KnowID]++
		if !one.NextLearnTime.Equal(theLastDay) && one.NextLearnTime.Before(now) {
			learning[one.KnowID]++
		}
	}

	k.valueCount(all, learning)
	return nil
}

func (k *KnowNode) valueCount(all, learning map[int64]int) {
	for _, one := range k.Children {
		one.valueCount(all, learning)
	}

	for _, one := range k.Children {
		k.TotalCount += all[one.ID]
		k.LearningCount += learning[one.ID]
	}
	k.TotalCount += all[k.ID]
	k.LearningCount += learning[k.ID]
}

func (k *know) GetOneByID(ctx shared.Context, learner *Learner, id int64) (*KnowNode, error) {
	where := &dao.KnowPointParam{
		Id: &id,
	}
	if learner != nil {
		where.Uid = &learner.ID
	}
	return k.getOne(ctx, where)
}

func (k *know) GetOneByName(ctx shared.Context, learner *Learner, name string) (*KnowNode, error) {
	where := &dao.KnowPointParam{
		Name: &name,
	}
	if learner != nil {
		where.Uid = &learner.ID
	}
	return k.getOne(ctx, where)
}

func (k *know) getOne(ctx shared.Context, where *dao.KnowPointParam) (*KnowNode, error) {
	know := &dao.KnowPoint{}
	err := know.Query(ctx, dao.XimuFlashcard(), where, nil)
	if err != nil {
		if shared.EmptyResult(err) {
			return nil, nil
		}
		ctx.Error(err)
		return nil, err
	}

	return &KnowNode{
		Label: know.Name,
		ID:    know.Id,
		Index: know.Index,
		PID:   know.Pid,
	}, nil
}

func (k *know) Update(ctx shared.Context, learner *Learner, id int64, param *dao.KnowPointParam) (int64, error) {
	rst, err := param.Update(ctx, dao.XimuFlashcard(), &dao.KnowPointParam{
		Id:  &id,
		Uid: &learner.ID,
	})

	if err != nil {
		ctx.Error(err)
		return 0, err
	}

	rowsAffected, err := rst.RowsAffected()
	if err != nil {
		ctx.Error(err)
		return 0, err
	}
	return rowsAffected, nil
}

func (k *know) Create(ctx shared.Context, learner *Learner, parentID int64, name string) (*KnowNode, error) {
	kno := &dao.KnowPointParam{
		Pid:  &parentID,
		Name: &name,
		Uid:  &learner.ID,
	}

	count, err := kno.MaxIndex(ctx, dao.XimuFlashcard())
	if err != nil {
		ctx.Error(err)
		return nil, err
	}
	count++
	kno.Index = &count

	rst, err := kno.Create(ctx, dao.XimuFlashcard())
	if err != nil {
		ctx.Error(err)
		return nil, err
	}
	id, err := rst.LastInsertId()
	if err != nil {
		ctx.Error(err)
		return nil, err
	}

	return &KnowNode{
		ID:    id,
		Label: name,
	}, nil
}

func (k *know) Delete(ctx shared.Context, learner *Learner, id int64) (int64, error) {
	kno := &dao.KnowPointParam{
		Id:  &id,
		Uid: &learner.ID,
	}

	rst, err := kno.Delete(ctx, dao.XimuFlashcard())
	if err != nil {
		ctx.Error(err)
		return 0, err
	}

	rows, err := rst.RowsAffected()
	if err != nil {
		ctx.Error(err)
		return 0, err
	}

	return rows, nil
}

func (k *know) Drap(ctx shared.Context, learner *Learner, id, rid int64, typ string) (err error) {
	relateNode, err := k.GetOneByID(ctx, learner, rid)
	if err != nil {
		return err
	}
	if relateNode == nil {
		ctx.Infof("bad rid: %v", rid)
		return nil
	}

	db, err := dao.XimuFlashcard().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			err = db.Commit()
			return
		}

		if err = db.Rollback(); err != nil {
			ctx.Error(err)
		}
	}()

	where := &dao.KnowPointParam{
		Id: &id,
	}

	node := &dao.KnowPointParam{}
	if typ == "inner" {
		var indexZero int64
		node.Pid = &rid
		node.Index = &indexZero

		if err = dao.KnowPointUpdateIndex(ctx, db, rid, 0); err != nil {
			ctx.Error(err)
			return
		}
	} else {
		index := relateNode.Index
		if typ == "after" {
			index++
		}
		if err = dao.KnowPointUpdateIndex(ctx, db, relateNode.PID, index); err != nil {
			ctx.Error(err)
			return
		}
		node.Index = &index
		node.Pid = &relateNode.PID
	}

	if _, err = node.Update(ctx, db, where); err != nil {
		ctx.Error(err)
		return
	}

	return nil
}
