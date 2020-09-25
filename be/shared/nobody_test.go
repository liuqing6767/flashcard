package shared

import "testing"

func TestIsEnglishWord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "1",
			s:    "hello",
			want: true,
		},
		{
			name: "2",
			s:    "hello word",
			want: false,
		},
		{
			name: "3",
			s:    "hello 中国",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEnglishWord(tt.s); got != tt.want {
				t.Errorf("IsEnglishWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
