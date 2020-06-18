package shared

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	TokenTypeRegister int8 = 1
	TokenTypeLogin    int8 = 2
)

type Token struct {
	ID         int64
	CreateTime time.Time
	Type       int8
}

func NewToken(token string) (*Token, error) {
	bs, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	plaintext, err := AESDecrypt(bs, AesKey)
	if err != nil {
		return nil, err
	}

	text := string(plaintext)
	group := strings.Split(text, ".")
	if len(group) != 3 {
		return nil, fmt.Errorf("bad plaintext: %s", text)
	}

	id, err := strconv.Atoi(group[1])
	if err != nil {
		return nil, fmt.Errorf("bad plaintext: %s", text)
	}
	createTime, err := strconv.Atoi(group[2])
	if err != nil {
		return nil, fmt.Errorf("bad plaintext: %s", text)
	}

	return &Token{
		ID:         int64(id),
		CreateTime: time.Unix(int64(createTime), 0),
	}, nil
}

func (l *Token) String() (string, error) {
	if l.CreateTime.IsZero() {
		l.CreateTime = time.Now()
	}
	plaintext := fmt.Sprintf("%d.%d.%d", l.Type, l.ID, l.CreateTime.Unix())
	bs, err := AESEncrypt([]byte(plaintext), AesKey)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bs), nil
}