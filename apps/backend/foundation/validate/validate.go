// Package validate contains the support for validating models.
package validate

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// Validate holds the settings and caches for validating request struct values.
var validate *validator.Validate

// Translator is a cache of locale and translation information.
var translator ut.Translator

func init() {
	validate = validator.New()

	// Create a translator for english so the error messages are
	// more human-readable than technical.
	translator, _ = ut.New(en.New(), en.New()).GetTranslator("en")

	// Register the english error messages for use.
	en_translations.RegisterDefaultTranslations(validate, translator)

	// Use JSON tag names for errors instead of Go struct names.
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

// Check validates the provided model against it's declared tags.
func Check(val any) error {
	if err := validate.Struct(val); err != nil {
		vErrors, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}

		fields := make(FieldErrors, len(vErrors))
		for i, vError := range vErrors {
			field := FieldError{Field: vError.Field(), Err: vError.Translate(translator)}
			fields[i] = field
		}

		return fields
	}
	return nil
}
