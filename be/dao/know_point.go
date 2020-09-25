// AUTO GEN BY dao, Modify it as you want
package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/liuximu/sqlmy"
)

var (
	DeletedNo  int8 = 0
	DeletedYes int8 = 1
)

const knowPointTable = "know_point"

type KnowPoint struct {
	Id         int64     `ddb:"id" json:"id"`
	Uid        int64     `ddb:"uid" json:"uid"`
	Pid        int64     `ddb:"pid" json:"pid"`
	Index      int64     `ddb:"inde" json:"inde"`
	Name       string    `ddb:"name" json:"name"`
	Status     int8      `ddb:"status" json:"status"`
	CreateTime time.Time `ddb:"create_time" json:"create_time"`
}

type KnowPointParam struct {
	Id         *int64     `ddb:"id" json:"id"`
	Uid        *int64     `ddb:"uid" json:"uid"`
	Pid        *int64     `ddb:"pid" json:"pid"`
	Index      *int64     `ddb:"inde" json:"inde"`
	Name       *string    `ddb:"name" json:"name"`
	Status     *int8      `ddb:"status" json:"status"`
	OrderBy    *string    `ddb:"_orderby"`
	CreateTime *time.Time `ddb:"create_time" json:"create_time"`
}

type KnowPointList []*KnowPoint
type KnowPointParamList []*KnowPoint

func (kp *KnowPoint) Query(ctx context.Context, db sqlmy.QueryExecer, where *KnowPointParam, columns []string) error {
	return db.QueryRowContextWithBuilderStruct(ctx, sqlmy.NewSelectBuilder(knowPointTable, sqlmy.Struct2Where(where), columns), kp)
}

func (kpp *KnowPointParam) Create(ctx context.Context, db sqlmy.QueryExecer) (sql.Result, error) {
	return db.ExecContextWithBuilder(ctx, sqlmy.NewInsertBuilder(knowPointTable, sqlmy.Struct2AssignList(kpp)))
}

func (kpp *KnowPointParam) MaxIndex(ctx context.Context, db sqlmy.QueryExecer) (count int64, err error) {
	sql := fmt.Sprintf(`SELECT COALESCE(MAX(inde), 0) FROM %s WHERE pid = ?`, knowPointTable)
	r := db.QueryRowContext(ctx, sql, *kpp.Pid)

	err = r.Scan(&count)
	return count, err
}

func (kpp *KnowPointParam) Update(ctx context.Context, db sqlmy.QueryExecer, where *KnowPointParam) (sql.Result, error) {
	return db.ExecContextWithBuilder(ctx, sqlmy.NewUpdateBuilder(knowPointTable, sqlmy.Struct2Where(where), sqlmy.Struct2Assign(kpp)))
}

func (kpp *KnowPointParam) Delete(ctx context.Context, db sqlmy.QueryExecer) (sql.Result, error) {
	return db.ExecContextWithBuilder(ctx, sqlmy.NewDeleteBuilder(knowPointTable, sqlmy.Struct2Where(kpp)))
}

func (kpl *KnowPointList) Query(ctx context.Context, db sqlmy.QueryExecer, where *KnowPointParam, columns []string) error {
	fmt.Println(sqlmy.Struct2Where(where))
	return db.QueryContextWithBuilderStruct(ctx, sqlmy.NewSelectBuilder(knowPointTable, sqlmy.Struct2Where(where), columns), kpl)
}

func (kppl KnowPointParamList) Create(ctx context.Context, db sqlmy.QueryExecer) (sql.Result, error) {
	_kppl := make([]interface{}, len(kppl))
	for i, one := range kppl {
		_kppl[i] = one
	}
	return db.ExecContextWithBuilder(ctx, sqlmy.NewInsertBuilder(knowPointTable, sqlmy.Struct2AssignList(_kppl...)))
}

func KnowPointUpdateIndex(ctx context.Context, db sqlmy.QueryExecer, pid, startIndex int64) error {
	sql := fmt.Sprintf("UPDATE %s SET inde = inde+1 WHERE pid = ? AND inde >= ?", knowPointTable)
	_, err := db.ExecContext(ctx, sql, pid, startIndex)
	return err
}
