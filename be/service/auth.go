package service

import (
	"fmt"
	"time"

	"github.com/liuximu/flashcard/be/dao"
	"github.com/liuximu/flashcard/be/shared"
)

type auth struct{}

var Auth = &auth{}

type Learner struct {
	ID           int64
	Email        string
	Status       int8
	password     string
	RegisterTime time.Time

	Token *shared.Token
}

var (
	SendEmailSpan = 5 * time.Minute
)

func (l *Learner) GenToken(typ int8) (string, error) {
	token := &shared.Token{
		ID:   l.ID,
		Type: typ,
	}

	return token.String()
}

func (a *auth) GetUserByToken(ctx shared.Context, tokens string) (*Learner, error) {
	token, err := shared.NewToken(tokens)
	if err != nil {
		ctx.Error(err)
		return nil, err
	}
	learner, err := a.getUser(ctx, &dao.AuthLearnerParam{
		Id: &token.ID,
	})

	if err != nil {
		return nil, err
	}

	learner.Token = token
	return learner, nil
}

func (l *Learner) RightPassword(pw string) bool {
	return l.password == pw
}

func (a *auth) getUser(ctx shared.Context, where *dao.AuthLearnerParam) (*Learner, error) {
	al := &dao.AuthLearner{}
	err := al.Query(ctx, dao.XimuFlashcard(), where, nil)
	if err != nil {
		if shared.EmptyResult(err) {
			return nil, nil
		}
		ctx.Errorf("GetUserByEmail fail, err: %v", err)
		return nil, err
	}

	learner := &Learner{
		ID:           al.Id,
		Email:        al.Email,
		Status:       al.Status,
		password:     al.Password,
		RegisterTime: al.RegisterTime,
	}

	return learner, nil
}

func (a *auth) GetUserByEmail(ctx shared.Context, email string) (*Learner, error) {

	return a.getUser(ctx, &dao.AuthLearnerParam{
		Email: &email,
	})
}

func (a *auth) CreateUser(ctx shared.Context, email, password string) (*Learner, error) {
	al := &dao.AuthLearnerParam{
		Email:    &email,
		Password: &password,
		Status:   &dao.LearnerStatusRegisting,
	}
	_, err := al.Create(ctx, dao.XimuFlashcard())
	if err != nil {
		ctx.Errorf("CreateUser fail, err: %v", err)
		return nil, err
	}

	return a.GetUserByEmail(ctx, email)
}

func (a *auth) UpdateUser(ctx shared.Context, email string, learnerParam *dao.AuthLearnerParam) (*Learner, error) {
	_, err := learnerParam.Update(ctx, dao.XimuFlashcard(), &dao.AuthLearnerParam{
		Email: &email,
	})
	if err != nil {
		ctx.Errorf("UpdateUser fail, err: %v", err)
		return nil, err
	}

	return a.GetUserByEmail(ctx, email)
}

func (l *Learner) SendRegisterEmail(ctx shared.Context) error {
	token, err := l.GenToken(shared.TokenTypeRegister)
	if err != nil {
		ctx.Errorf("SendRegisterEmail fail, err: %v", err)
		return err
	}
	link := fmt.Sprintf("%s?token=%s", shared.AppConfSingleton.Links.RegisterActive, token)

	if err := shared.EmailSenderSingleton.Send(shared.Email{
		IsHtml:     true,
		Subject:    "欢迎注册闪记卡",
		Recipients: []string{l.Email},
		Content: fmt.Sprintf(`
<html><body>

Hello, Dear: </br>
欢迎使用闪记卡,更好的维护自己的知识体系。</br>

请点击 <a href="%s">激活账号</a> 完成激活。 </br>
如果无法点击，请复制链接在浏览器中完成激活： %s </br></br>
自动邮件，请勿回复。 </br>
	
闪记卡·熙穆

</body></html> 
		`, link, link),
	}); err != nil {
		ctx.Errorf("SendRegisterEmail fail, err: %v", err)
		return err
	}

	return nil
}
