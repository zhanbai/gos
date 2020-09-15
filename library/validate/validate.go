package validate

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Get(br interface{}) (bool, string) {
	zh := zh.New()
	uni = ut.New(zh)
	trans, _ := uni.GetTranslator("zh")

	validate = validator.New()
	zh_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(br)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return true, err.Error()
		}

		for _, err := range err.(validator.ValidationErrors) {
			return true, err.Translate(trans)
		}
	}

	return false, ""
}
