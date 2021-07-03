package interactor

import (
	"api/domain/model"
	"api/usecase/mock_repository"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_registerUserInteractor_RegisterUser(t *testing.T) {

	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		r func(*mock_repository.MockRegisterUserRepository)
	}
	type args struct {
		req model.RegisterUserRequestParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				r: func(mrur *mock_repository.MockRegisterUserRepository) {
					mrur.EXPECT().
						InsertData(gomock.Any()).
						Return(nil)
				},
			},
			args: args{
				req: model.RegisterUserRequestParam{
					FirstName:       "sample",
					LastName:        "sample",
					Password:        "password",
					ConfirmPassword: "password",
				},
			},
			wantErr: false,
		},
		{
			name: "password mismatch",
			fields: fields{
				r: func(mrur *mock_repository.MockRegisterUserRepository) {},
			},
			args: args{
				req: model.RegisterUserRequestParam{
					FirstName:       "sample",
					LastName:        "sample",
					Password:        "password",
					ConfirmPassword: "differentPassword",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := mock_repository.NewMockRegisterUserRepository(ctrl)
			tt.fields.r(m)

			i := &registerUserInteractor{
				r: m,
			}
			if err := i.RegisterUser(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("registerUserInteractor.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
