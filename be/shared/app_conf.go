package shared

import (
	"io/ioutil"

	"encoding/json"

	"github.com/labstack/echo/v4"
	"github.com/liuximu/flashcard/be/dao"
	"github.com/liuximu/sqlmy"
)

type AppConf struct {
    Port     string
	TokenKey string
	Links    struct {
		RegisterActive string
		LoginPage      string
	}
	Email struct {
		User     string
		Password string
		Host     string
	}
	XimuDB *sqlmy.MyDB
}

func LoadAppConf(file string, log echo.Logger) error {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	conf := &AppConf{}
	if err := json.Unmarshal(bs, conf); err != nil {
		return err
	}

	AppConfSingleton = conf
	EmailSenderSingleton = &EmailSender{
		User:     conf.Email.User,
		Password: conf.Email.Password,
		Host:     conf.Email.Host,
	}

	ximuDB := conf.XimuDB
	ximuDB.StatusLogger = log
	ximuDB.AccessLog = log
	ximuDB.Hooks = sqlmy.NewLogHooks()
	if err := ximuDB.Init(); err != nil {
		return err
	}
	dao.SetXimuFlashcard(ximuDB)

	AesKey = []byte(conf.TokenKey)

	return nil
}

var (
	AppConfSingleton     *AppConf
	EmailSenderSingleton *EmailSender
	AesKey               []byte

	// XimuDBSingleton *sqlmy.MyDB
)
