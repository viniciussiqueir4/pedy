package models

import (
	"errors"
	"fmt"
	en2 "github.com/go-playground/locales/en"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_BR2 "github.com/go-playground/validator/v10/translations/pt_BR"
)
// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func SetValidationPtBr() (error, ut.Translator) {
	pt := pt_BR.New()
	en := en2.New()
	uni = ut.New(en, pt)

	trans, found := uni.GetTranslator("pt_BR")
	if !found {
		return errors.New("Could not find desired translation pt_BR"), nil
	}
	Validate = validator.New()
	err := pt_BR2.RegisterDefaultTranslations(Validate, trans)
	if err != nil {
		return err, nil
	}
	return nil, trans
}

func TranslateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}