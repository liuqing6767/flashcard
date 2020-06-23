// AUTO GEN BY dao, Modify it as you want
package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/liuximu/sqlmy"
)

const cardTable = "card"

type Card struct {
	Id         int64  `ddb:"id" json:"id"`
	Name       string `ddb:"name" json:"name"`
	KnowID     int64  `ddb:"kid" json:"kid"`
	TemplateID int64  `ddb:"tid" json:"tid"`
	Uid        int64  `ddb:"uid" json:"uid"`
	Data       string `ddb:"data" json:"data"`
	Status     int8   `ddb:"status" json:"status"`

	MemoryCurveID int64     `ddb:"mid" json:"mid"`
	Progress      int8      `ddb:"progress" json:"progress"`
	NextLearnTime time.Time `ddb:"next_learn_time" json:"next_learn_time"`

	CreateTime time.Time `ddb:"create_time" json:"create_time"`
}

type CardParam struct {
	Id         *int64  `ddb:"id" json:"id"`
	Name       *string `ddb:"name" json:"name"`
	KnowID     *int64  `ddb:"kid" json:"kid"`
	TemplateID *int64  `ddb:"tid" json:"tid"`
	Uid        *int64  `ddb:"uid" json:"uid"`
	Data       *string `ddb:"data" json:"data"`
	Status     *int8   `ddb:"status" json:"status"`

	MemoryCurveID    *int64     `ddb:"mid" json:"mid"`
	Progress         *int8      `ddb:"progress" json:"progress"`
	NextLearnTime    *time.Time `ddb:"next_learn_time" json:"next_learn_time"`
	MaxNextLearnTime *time.Time `ddb:"next_learn_time,<"`

	CreateTime *time.Time `ddb:"create_time" json:"create_time"`

	OrderBy *string `ddb:"_orderby"`
}

type CardList []*Card
type CardParamList []*CardParam

func (c *Card) Query(ctx context.Context, db sqlmy.QueryExecer, where *CardParam, columns []string) error {
	return db.QueryRowContextWithBuilderStruct(ctx, sqlmy.NewSelectBuilder(cardTable, sqlmy.Struct2Where(where), columns), c)
}

func (cp *CardParam) Create(ctx context.Context, db sqlmy.QueryExecer) (sql.Result, error) {
	return db.ExecContextWithBuilder(ctx, sqlmy.NewInsertBuilder(cardTable, sqlmy.Struct2AssignList(cp)))
}

func (cp *CardParam) Update(ctx context.Context, db sqlmy.QueryExecer, where *CardParam) (sql.Result, error) {
	return db.ExecContextWithBuilder(ctx, sqlmy.NewUpdateBuilder(cardTable, sqlmy.Struct2Where(where), sqlmy.Struct2Assign(cp)))
}

func (cp *CardParam) Delete(ctx context.Context, db sqlmy.QueryExecer) (sql.Result, error) {
	return db.ExecContextWithBuilder(ctx, sqlmy.NewDeleteBuilder(cardTable, sqlmy.Struct2Where(cp)))
}

func (cl *CardList) Query(ctx context.Context, db sqlmy.QueryExecer, where *CardParam, columns []string) error {
	return db.QueryContextWithBuilderStruct(ctx, sqlmy.NewSelectBuilder(cardTable, sqlmy.Struct2Where(where), columns), cl)
}

func (cpl CardParamList) Create(ctx context.Context, db sqlmy.QueryExecer) (sql.Result, error) {
	_cpl := make([]interface{}, len(cpl))
	for i, one := range cpl {
		_cpl[i] = one
	}
	return db.ExecContextWithBuilder(ctx, sqlmy.NewInsertBuilder(cardTable, sqlmy.Struct2AssignList(_cpl...)))
}

func CardCountLearning(ctx context.Context, db sqlmy.QueryExecer, uid int64, tim time.Time) (int64, error) {
	row := db.QueryRowContext(ctx, "SELECT count(1) FROM card WHERE uid = ? AND next_learn_time < ?", uid, tim)
	var count int64
	err := row.Scan(&count)
	return count, err
}
