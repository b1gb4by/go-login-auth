package testdata

import "api/domain/model"

var LoginAuthenticationRequestParam = map[string]*model.LoginAuthenticationRequestParam{

	"success_1": {
		Email:    "sample@example.com",
		Password: "sample123456",
	},

	"failed_1": {
		Email:    "sample@example.com",
		Password: "sample654321",
	},
}
