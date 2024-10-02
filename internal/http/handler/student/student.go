package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kundannishad/students-api/internal/types/student"
	"github.com/kundannishad/students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Create A student api")

		var student student.StudentStr

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			//w.Write([]byte("Welcome to student api"))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))

			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		}

		//Request Validation

		//err := validate.Struct(&student)
		//validationErrors := err.(validator.ValidationErrors)

		if err := validator.New().Struct(student); err != nil {
			validateError := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateError))
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "Heloo"})

	}
}
