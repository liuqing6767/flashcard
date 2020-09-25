// AUTO GEN BY dao, Modify it as you want
package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/liuximu/sqlmy"
)

const authLearnerTable = "auth_learner"

type AuthLearner struct {
	Id           int64     `ddb:"id"`
	Email        string    `ddb:"email"`
	Password     string    `ddb:"password"`
	RegisterTime time.Time `ddb:"register_time"`
	Status       int8      `ddb:"status"`
}

type AuthLearnerParam struct {
	Id           *int64     `ddb:"id"`
	Email        *string    `ddb:"email"`
	Password     *string    `ddb:"password"`
	RegisterTime *time.Time `ddb:"register_time"`
	Status       *int8      `ddb:"status"`
}

var (
	LearnerStatusRegisting int8 = 1
	LearnerStatuOk         int8 = 2
)

type AuthLearnerList []*AuthLearner
type AuthLearnerParamList []*AuthLearner

func (al *AuthLearner) Query(ctx context.Context, db sqlmy.QueryExecer, where *AuthLearnerParam, columns []string) error {
	return db.QueryRowContextWithBuilderStruct(ctx, sqlmy.NewSelectBuilder(authLearnerTable, sqlmy.Struct2Where(where), columns), al)
}

func (alp *AuthLearnerParam) Create(ctx context.Context, db sqlmy.QueryExecer) (sql.Result, error) {
	return db.ExecContextWithBuilder(ctx, sqlmy.NewInsertBuilder(authLearnerTable, sqlmy.Struct2AssignList(alp)))
}

func (alp *AuthLearnerParam) Update(ctx context.Context, db sqlmy.QueryExecer, where *AuthLearnerParam) (sql.Result, error) {
	return db.ExecContextWithBuilder(ctx, sqlmy.NewUpdateBuilder(authLearnerTable, sqlmy.Struct2Where(where), sqlmy.Struct2Assign(alp)))
}

func (all *AuthLearnerList) Query(ctx context.Context, db sqlmy.QueryExecer, where *AuthLearnerParam, columns []string) error {
	return db.QueryContextWithBuilderStruct(ctx, sqlmy.NewSelectBuilder(authLearnerTable, sqlmy.Struct2Where(where), columns), all)
}

func (alpl AuthLearnerParamList) Create(ctx context.Context, db sqlmy.QueryExecer) (sql.Result, error) {
	_alpl := make([]interface{}, len(alpl))
	for i, one := range alpl {
		_alpl[i] = one
	}
	return db.ExecContextWithBuilder(ctx, sqlmy.NewInsertBuilder(authLearnerTable, sqlmy.Struct2AssignList(_alpl...)))
}
