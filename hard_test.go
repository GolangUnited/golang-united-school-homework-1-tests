package hard

import (
	"strings"
	"testing"
)

func TestGoalParsers(t *testing.T) {
	type args struct {
		strReader *strings.Reader
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				strings.NewReader("G()(al)"),
			},
			want: "Goal",
		},
		{
			name: "",
			args: args{
				strings.NewReader("G()()()()(al)"),
			},
			want: "Gooooal",
		},
		{
			name: "",
			args: args{
				strings.NewReader("(al)G(al)()()G"),
			},
			want: "alGalooG",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GoalParsers(tt.args.strReader); got != tt.want {
				t.Errorf("GoalParsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
