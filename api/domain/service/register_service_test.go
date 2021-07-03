package service

import (
	"reflect"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestIsPasswordMatch(t *testing.T) {
	t.Parallel()
	type args struct {
		p  string
		cp string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Password match",
			args: args{
				p:  "test",
				cp: "test1",
			},
			wantErr: false,
		},
		{
			name: "Password mismatch",
			args: args{
				p:  "test",
				cp: "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := IsPasswordMatch(tt.args.p, tt.args.cp); (err != nil) != tt.wantErr {
				t.Errorf("IsPasswordMatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGeneratedFromPassword(t *testing.T) {
	t.Parallel()

	test1, _ := bcrypt.GenerateFromPassword([]byte("sample"), 14)

	type args struct {
		p string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Successful generation of hash values",
			args: args{
				p: "sample",
			},
			want:    test1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := GeneratedFromPassword(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratedFromPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GeneratedFromPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
