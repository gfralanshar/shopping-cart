package exception

import (
	"net/http"
	"shopping-chart/api/v1/helper"
	"shopping-chart/api/v1/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err any) {
	if validationErrors(w, r, err) {
		return
	}

	if notFoundError(w, r, err) {
		return
	}

	if permissionError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func validationErrors(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)

		webResponse := web.WebResponse{
			Status: "unporoccessable entity",
			Code:   http.StatusUnprocessableEntity,
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	}

	return false
}

func notFoundError(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Status: "not found",
			Code:   http.StatusNotFound,
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	}

	return false
}

func permissionError(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(PermissionError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)

		webResponse := web.WebResponse{
			Status: "forbidden error",
			Code: http.StatusForbidden,
			Data: exception.Error,
		}

		helper.WriteToResponseBody(w, webResponse)
		return true
	}

	return false
}

func internalServerError(w http.ResponseWriter, r *http.Request, err any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Status: "internal server error",
		Code:   http.StatusInternalServerError,
		Data:   err,
	}

	helper.WriteToResponseBody(w, webResponse)
}
