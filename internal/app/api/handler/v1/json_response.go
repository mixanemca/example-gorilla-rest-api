package v1

import (
	"encoding/json"
	"net/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/mixanemca/example-gorilla-rest-api/models"
)

// jsonResponse for convert response to json format
func jsonResponse(w http.ResponseWriter, status int, model interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if model != nil {
		jsonData, err := json.Marshal(model)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = w.Write(jsonData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// jsonErrResponse for convert error response to json format
func jsonErrResponse(w http.ResponseWriter, translator ut.Translator, err error) {
	var fieldsErrors []models.FieldError
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		fieldsErrors = append(fieldsErrors, models.FieldError{
			ErrorField:   e.Field(),
			ErrorMessage: e.Translate(translator),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	jsonData, err := json.Marshal(fieldsErrors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
