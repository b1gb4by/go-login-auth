package validation

import (
	"github.com/go-playground/validator/v10"
)

// RequestBodyValidate
// リクエストボディのバリデーションチェック処理.
func RequestBodyValidate(r interface{}) error {
	// Validateインスタンスを生成
	validate := validator.New()

	err := validate.Struct(r)

	return err
}
