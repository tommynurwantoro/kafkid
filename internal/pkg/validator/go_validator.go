package validator

import (
	"context"
	"errors"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type GoValidator struct {
	validate *validator.Validate
	uni      ut.Translator
}

func NewGoValidator() Validator {
	v := validator.New()
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(v, trans)

	return &GoValidator{validate: v, uni: trans}
}

func (v *GoValidator) Validate(ctx context.Context, data interface{}) error {
	err := v.validate.StructCtx(ctx, data)
	if err == nil {
		return nil
	}

	if _, ok := err.(*validator.InvalidValidationError); ok {
		return err
	}

	errs := err.(validator.ValidationErrors)
	if len(errs) > 0 {
		mapErr := make(map[string]error, 0)
		for _, err := range errs {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			mapErr[err.Field()] = errors.New(err.Translate(v.uni))
		}

		return NewErrorMap(mapErr)
	}

	return nil
}
