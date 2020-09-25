// AUTO GEN BY dao, Modify it as you want
package dao

import (
	"context"
	"database/sql"

	"github.com/liuximu/sqlmy"
)

const WordTemplateID int64 = 4

const cardTemplateTable = "card_template"

type CardTemplate struct {
	Id          int64  `ddb:"id" json:"id"`
	Name        string `ddb:"name" json:"name"`
	ContentA    string `ddb:"content_a" json:"content_a"`
	ContentB    string `ddb:"content_b" json:"content_b"`
	DataDemo    string `ddb:"data_demo" json:"data_demo"`
	CardNameKey string `ddb:"card_name_key" json:"card_name_key"`
	Uid         int64  `ddb:"uid" json:"uid"`
	Status      int8   `ddb:"status" json:"status"`
}

type CardTemplateParam struct {
	Id          *int64  `ddb:"id" json:"id"`
	Ids         []int64 `ddb:"id,in" json:"id"`
	Name        *string `ddb:"name" json:"name"`
	ContentA    *string `ddb:"content_a" json:"content_a"`
	ContentB    *string `ddb:"content_b" json:"content_b"`
	DataDemo    *string `ddb:"data_demo" json:"data_demo"`
	CardNameKey *string `ddb:"card_name_key" json:"card_name_key"`
	Uid         *int64  `ddb:"uid" json:"uid"`
	Uids        []int64 `ddb:"uid,in" json:"uid"`
	Status      *int8   `ddb:"status" json:"status"`
}

type CardTemplateList []*CardTemplate
type CardTemplateParamList []*CardTemplate

func (ct *CardTemplate) Query(ctx context.Context, db sqlmy.QueryExecer, where *CardTemplateParam, columns []string) error {
	return db.QueryRowContextWithBuilderStruct(ctx, sqlmy.NewSelectBuilder(cardTemplateTable, sqlmy.Struct2Where(where), columns), ct)
}

func (ctp *CardTemplateParam) Create(ctx context.Context, db sqlmy.QueryExecer) (sql.Result, error) {
	return db.ExecContextWithBuilder(ctx, sqlmy.NewInsertBuilder(cardTemplateTable, sqlmy.Struct2AssignList(ctp)))
}

func (ctp *CardTemplateParam) Update(ctx context.Context, db sqlmy.QueryExecer, where *CardTemplateParam) (sql.Result, error) {
	return db.ExecContextWithBuilder(ctx, sqlmy.NewUpdateBuilder(cardTemplateTable, sqlmy.Struct2Where(where), sqlmy.Struct2Assign(ctp)))
}

func (ctl *CardTemplateList) Query(ctx context.Context, db sqlmy.QueryExecer, where *CardTemplateParam, columns []string) error {
	return db.QueryContextWithBuilderStruct(ctx, sqlmy.NewSelectBuilder(cardTemplateTable, sqlmy.Struct2Where(where), columns), ctl)
}

func (ctpl CardTemplateParamList) Create(ctx context.Context, db sqlmy.QueryExecer) (sql.Result, error) {
	_ctpl := make([]interface{}, len(ctpl))
	for i, one := range ctpl {
		_ctpl[i] = one
	}
	return db.ExecContextWithBuilder(ctx, sqlmy.NewInsertBuilder(cardTemplateTable, sqlmy.Struct2AssignList(_ctpl...)))
}
