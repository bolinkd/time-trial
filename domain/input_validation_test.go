package domain

import (
	"testing"

	"gopkg.in/volatiletech/null.v6"
)

func Test_isValidNullInt(t *testing.T) {
	type args struct {
		x null.Int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Invalid Null Int64",
			args: args{
				null.Int{
					Int:   0,
					Valid: false,
				},
			},
			want: false,
		},
		{
			name: "Zero Null Int",
			args: args{
				null.Int{
					Int:   0,
					Valid: true,
				},
			},
			want: false,
		},
		{
			name: "Valid Null Int",
			args: args{
				null.Int{
					Int:   6,
					Valid: true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidNullInt(tt.args.x); got != tt.want {
				t.Errorf("isValidNullInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidNullInt64(t *testing.T) {

}

func Test_isValidNullString(t *testing.T) {

}

func Test_isValidNullTime(t *testing.T) {

}
