package shared

import (
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	tests := []struct {
		name    string
		token   *Token
		want    string
		wantErr bool
	}{
		{
			name: "case1",
			token: &Token{
				ID:         100,
				CreateTime: time.Unix(1111, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := tt.token.String()
			if (err != nil) != tt.wantErr {
				t.Errorf("Token.Token() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}

			tokenObj, err := NewToken(token)
			if err != nil {
				t.Error(err)
			}
			if tokenObj.ID != tt.token.ID ||
				tokenObj.CreateTime != tt.token.CreateTime {
				t.Errorf("%v, %v", tokenObj, tt.token)
			}
		})
	}
}
