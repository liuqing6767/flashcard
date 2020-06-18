package service

import (
	"time"

	"github.com/liuximu/flashcard/shared"
)

type memory struct{}

var Memory = &memory{}

type MemoryCurve struct {
	MaxStage int8
}

func (c *memory) GetByID(ctx shared.Context, rid int64) (*MemoryCurve, error) {
	return &MemoryCurve{
		MaxStage: 7,
	}, nil
}

// Next 下一个学习时间
// curStage [0, 7]
// familiarity [0,3]
func (mc *MemoryCurve) Next(curStage int8, familiarity int8) (nextStage int8, learnTime time.Time) {
	if curStage == mc.MaxStage {
		return curStage, theLastDay
	}

	duration, nextStage := shared.Next(curStage, familiarity)
	if nextStage == -1 {
		return mc.MaxStage, theLastDay

	}

	return nextStage, time.Now().Add(duration)
}

var theLastDay, _ = time.Parse("2006", "9999")
