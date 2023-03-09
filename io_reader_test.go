package homework_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestReaderSplit(t *testing.T) {
	type args struct {
		strReader *strings.Reader
		n         int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				strReader: strings.NewReader("helloworld"),
				n:         2,
			},
			want: []string{
				"he", "ll", "ow", "or", "ld",
			},
		},
		{
			name: "",
			args: args{
				strReader: strings.NewReader("GolangUnitedSchool"),
				n:         5,
			},
			want: []string{
				"Golan", "gUnit", "edSch", "ool",
			},
		},
		{
			name: "",
			args: args{
				strReader: strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
				n:         6,
			},
			want: []string{
				"ABCDEF", "GHIJKL", "MNOPQR", "STUVWX", "YZ",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := homework.ReaderSplit(tt.args.strReader, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got, tt.want)
				t.Errorf("ReaderSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadHalfOfString(t *testing.T) {
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
				strings.NewReader("HELLO-WORLD"),
			},
			want: "-WORLD",
		},
		{
			name: "",
			args: args{
				strings.NewReader("HELLOWORLD"),
			},
			want: "WORLD",
		},
		{
			name: "",
			args: args{
				strings.NewReader("GOLANG-UNITED"),
			},
			want: "-UNITED",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := homework.SeekTillHalfOfString(tt.args.strReader); got != tt.want {
				t.Errorf("SeekTillHalfOfString = %v, want %v", got, tt.want)
			}
		})
	}
}
