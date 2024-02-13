package middleware

import (
	"net/http"
	"shopping-chart/api/v1/helper"
	"shopping-chart/api/v1/security"
	"shopping-chart/api/v1/web"

	"github.com/julienschmidt/httprouter"
)

func AuthMiddleware(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		err := security.VerifyToken(r)
		if err != nil {
			helper.WriteToResponseBody(w, web.WebResponse{
				Status: "Unauthorized",
				Code:   http.StatusUnauthorized,
			})
			return
		}

		n(w, r, p)
	}
}
