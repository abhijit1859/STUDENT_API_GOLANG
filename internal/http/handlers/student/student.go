package student

import (
	"encoding/json"
	"errors"
	"fmt"

	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/proj/internal/storage"
	"github.com/proj/internal/types"
	"github.com/proj/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("creating a student")
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest,err.Error())
			return
		}

		if err!=nil{
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(err))
			return 
		}

		//validating request 
		//validating all the input fields 
		//we use validation-v10 library

		if err:=validator.New().Struct(student);err!=nil{
			
			validateErrs:=err.(validator.ValidationErrors)
			response.WriteJson(w,http.StatusBadRequest,response.ValidateErrors(validateErrs))
		}

		lastId,err:=storage.CreateStudent(student.Name,student.Email,student.Age)

 
		slog.Info("user created successfully", slog.String("userId", fmt.Sprint(lastId)))

		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}
