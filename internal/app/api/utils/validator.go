// Package utils for custom utilits like as custom validator and etc.
package utils

import (
	"regexp"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func NewValidator() (*validator.Validate, ut.Translator, error) {
	en := en.New()
	uni := ut.New(en, en)
	translate, _ := uni.GetTranslator("en")

	validate := validator.New()

	// set custom message for required validation
	err := validate.RegisterTranslation("required", translate, func(ut ut.Translator) error {
		return ut.Add("required", "Не заполнено обязательное поле", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
	if err != nil {
		return nil, nil, err
	}

	// set custom message for only alpha validation
	err = validate.RegisterTranslation("alpha", translate, func(ut ut.Translator) error {
		return ut.Add("alpha", "Может содержать только английские буквы", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("alpha", fe.Field())
		return t
	})
	if err != nil {
		return nil, nil, err
	}

	// set custom message for alpha and numbers validation
	err = validate.RegisterTranslation("alphanum", translate, func(ut ut.Translator) error {
		return ut.Add("alphanum", "Может содержать только английские буквы и цифры", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("alphanum", fe.Field())
		return t
	})
	if err != nil {
		return nil, nil, err
	}

	// set custom message for minimal symbols validation
	err = validate.RegisterTranslation("min", translate, func(ut ut.Translator) error {
		return ut.Add("min", "Необходимо указать более трех символов", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field())
		return t
	})
	if err != nil {
		return nil, nil, err
	}

	// set custom message for email validation
	err = validate.RegisterTranslation("email", translate, func(ut ut.Translator) error {
		return ut.Add("email", "Ошибка формата заполнения адреса почты", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})
	if err != nil {
		return nil, nil, err
	}

	// set custom message for phone validation
	err = validate.RegisterTranslation("e164", translate, func(ut ut.Translator) error {
		return ut.Add("e164", "Ошибка формата заполнения номера телефона", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("e164", fe.Field())
		return t
	})
	if err != nil {
		return nil, nil, err
	}

	return validate, translate, nil
}

func ValidUUID(id string) bool {
	re := regexp.MustCompile(`[a-fA-F\d]{8}-[a-fA-F\d]{4}-[a-fA-F\d]{4}-[a-fA-F\d]{4}-[a-fA-F\d]{12}$`)
	return re.MatchString(id)
}
