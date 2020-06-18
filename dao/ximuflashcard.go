// AUTO GEN BY  dao, Modify it as you want
package dao

import (
	"context"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/liuximu/sqlmy"
)

var ximuFlashcardSingleton sqlmy.DB

func SetXimuFlashcard(ximuFlashcard sqlmy.DB) {
	ximuFlashcardSingleton = ximuFlashcard
}

func XimuFlashcard() sqlmy.DB {
	return ximuFlashcardSingleton
}

func MockXimuFlashcard() sqlmock.Sqlmock {
	db, mock, err := sqlmy.MockDb()
	if err != nil {
		panic(err)
	}

	ximuFlashcardSingleton = db
	return mock
}

func NewXimuFlashcardTx(ctx context.Context) (sqlmy.Tx, error) {
	return XimuFlashcard().BeginTx(ctx, nil)
}
