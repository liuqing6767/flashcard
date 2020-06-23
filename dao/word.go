// AUTO GEN BY dao, Modify it as you want
package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/liuximu/sqlmy"
)

var (
	WordStatusNoExisted int8 = 1
)

const wordTable = "word"

type Word struct {
	Id         int32     `ddb:"id" json:"id"`
	Word       string    `ddb:"word" json:"word"`
	Content    string    `ddb:"content" json:"content"`
	Status     int8      `ddb:"status" json:"status"`
	CreateTime time.Time `ddb:"create_time" json:"create_time"`
}

type WordParam struct {
	Id         *int32     `ddb:"id" json:"id"`
	Word       *string    `ddb:"word" json:"word"`
	Content    *string    `ddb:"content" json:"content"`
	Status     *int8      `ddb:"status" json:"status"`
	CreateTime *time.Time `ddb:"create_time" json:"create_time"`
}

type WordList []*Word
type WordParamList []*WordParam

func (w *Word) Query(ctx context.Context, db sqlmy.QueryExecer, where *WordParam, columns []string) error {
	return db.QueryRowContextWithBuilderStruct(ctx, sqlmy.NewSelectBuilder(wordTable, sqlmy.Struct2Where(where), columns), w)
}

func (wp *WordParam) Create(ctx context.Context, db sqlmy.QueryExecer) (sql.Result, error) {
	return db.ExecContextWithBuilder(ctx, sqlmy.NewInsertBuilder(wordTable, sqlmy.Struct2AssignList(wp)))
}

func (wp *WordParam) Update(ctx context.Context, db sqlmy.QueryExecer, where *WordParam) (sql.Result, error) {
	return db.ExecContextWithBuilder(ctx, sqlmy.NewUpdateBuilder(wordTable, sqlmy.Struct2Where(where), sqlmy.Struct2Assign(wp)))
}

func (wp *WordParam) Delete(ctx context.Context, db sqlmy.QueryExecer) (sql.Result, error) {
	return db.ExecContextWithBuilder(ctx, sqlmy.NewDeleteBuilder(wordTable, sqlmy.Struct2Where(wp)))
}

func (wl *WordList) Query(ctx context.Context, db sqlmy.QueryExecer, where *WordParam, columns []string) error {
	return db.QueryContextWithBuilderStruct(ctx, sqlmy.NewSelectBuilder(wordTable, sqlmy.Struct2Where(where), columns), wl)
}

func (wpl WordParamList) Create(ctx context.Context, db sqlmy.QueryExecer) (sql.Result, error) {
	_wpl := make([]interface{}, len(wpl))
	for i, one := range wpl {
		_wpl[i] = one
	}
	return db.ExecContextWithBuilder(ctx, sqlmy.NewInsertBuilder(wordTable, sqlmy.Struct2AssignList(_wpl...)))
}
