package service

import (
	"encoding/json"

	"github.com/liuximu/flashcard/dao"
	"github.com/liuximu/flashcard/shared"
)

type word struct{}

var Word = &word{}

func (w *word) QueryOrCreate(ctx shared.Context, wor string) (*shared.Word, error) {
	wordDao := &dao.Word{}
	err := wordDao.Query(ctx, dao.XimuFlashcard(), &dao.WordParam{Word: &wor}, nil)
	if err != nil {
		if !shared.EmptyResult(err) {
			ctx.Error(err)
			return nil, err
		}
		wordDao = nil
	}

	if wordDao != nil {
		if wordDao.Status == dao.WordStatusNoExisted {
			return nil, nil
		}

		wo := &shared.Word{}
		if err := json.Unmarshal([]byte(wordDao.Content), wo); err != nil {
			ctx.Error(err)
			return nil, err
		}
		return wo, nil
	}

	wo, err := shared.NewWord(wor)
	if err != nil {
		ctx.Error(err)
		return nil, err
	}

	bs, err := json.Marshal(wo)
	if err != nil {
		ctx.Error(err)
		return wo, nil
	}

	content := ""
	wp := &dao.WordParam{
		Word: &wor,
	}
	if wo.Def == "" {
		wp.Status = &dao.WordStatusNoExisted
	} else {
		content = string(bs)
	}
	wp.Content = &content

	_, err = wp.Create(ctx, dao.XimuFlashcard())
	if err != nil {
		ctx.Error(err)
	}

	return wo, nil
}
