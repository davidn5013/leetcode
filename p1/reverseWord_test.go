package p1

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Exempel 1",
			args: args{s: "Let's take LeetCode contest"},
			want: "s'teL ekat edoCteeL tsetnoc",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = \n[%v]\n, want \n[%v]\n", got, tt.want)
			}
		})
	}
}

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reverse a word",
			args: args{s: "word"},
			want: "dorw",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := reverseStr(tt.args.s); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
