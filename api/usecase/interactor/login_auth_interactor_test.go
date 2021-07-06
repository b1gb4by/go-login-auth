package interactor

import (
	"api/config"
	"api/domain/model"
	"api/testdata"
	"api/usecase/mock_repository"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_loginAuthenticationInteractor_LoginAuthentication(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		r  func(*mock_repository.MockLoginAuthenticationRepository)
		jc *config.JWTConfig
	}
	type args struct {
		req model.LoginAuthenticationRequestParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.AcquisitionUser
		want1   string
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				r: func(mlar *mock_repository.MockLoginAuthenticationRepository) {
					mlar.EXPECT().
						SearchUser("sample@example.com").
						Return(*testdata.AcquisitionUser["success_1"], nil)
				},
				jc: &config.JWTConfig{
					Secret: "sample_secret",
				},
			},
			args: args{
				req: *testdata.LoginAuthenticationRequestParam["success_1"],
			},
			want:    *testdata.AcquisitionUser["success_1"],
			wantErr: false,
		},
		{
			name: "Failed",
			fields: fields{
				r: func(mlar *mock_repository.MockLoginAuthenticationRepository) {
					mlar.EXPECT().
						SearchUser("sample@example.com").
						Return(model.AcquisitionUser{}, nil)
				},
				jc: &config.JWTConfig{
					Secret: "sample_secret",
				},
			},
			args: args{
				req: *testdata.LoginAuthenticationRequestParam["failed_1"],
			},
			want:    model.AcquisitionUser{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := mock_repository.NewMockLoginAuthenticationRepository(ctrl)
			tt.fields.r(m)

			i := &loginAuthenticationInteractor{
				r:  m,
				jc: tt.fields.jc,
			}
			got, _, err := i.LoginAuthentication(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginAuthentication() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoginAuthentication() got = %v, want %v", got, tt.want)
			}
		})
	}
}
