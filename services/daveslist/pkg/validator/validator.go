package validator

import (
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	pv "github.com/go-playground/validator/v10"
)

type Validator interface {
	StrcutWithTranslateError(s interface{}) []Error
}

type validator struct {
	*pv.Validate
	trans ut.Translator
}

func New(trans ut.Translator) Validator {
	v := pv.New()
	err := RegisterDefaultTranslations(v, trans)
	if err != nil {
		log.Fatal("validator: ", err)
	}
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := ""
		tagJSON := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		tagHeader := strings.SplitN(fld.Tag.Get("header"), ",", 2)[0]
		tagQuery := strings.SplitN(fld.Tag.Get("query"), ",", 2)[0]
		if tagJSON != "" {
			name = strings.ReplaceAll(tagJSON, "-", "")
		} else if tagQuery != "" {
			name = strings.ReplaceAll(tagQuery, "-", "")
		} else if tagHeader != "" {
			name = strings.ReplaceAll(tagHeader, "-", "")
		}
		return name
	})
	return &validator{v, trans}
}

func NewTranslator() ut.Translator {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	return trans
}

func (v *validator) StrcutWithTranslateError(s interface{}) []Error {
	err := v.Struct(s)
	if err == nil {
		return nil
	}
	validatorErrs := err.(pv.ValidationErrors)
	var errs []Error
	for _, e := range validatorErrs {
		errs = append(errs, NewError(e.Translate(v.trans)))
	}
	return errs
}
