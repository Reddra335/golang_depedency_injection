package middleware

import (
	"net/http"
	"restful_api/helper"
	"restful_api/web"
)

type AuthMidleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMidleware {

	return &AuthMidleware{
		Handler: handler,
	}

}
func (middleware *AuthMidleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if "RAHASIA" == request.Header.Get("X-API-Key") {
		//ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {

		//salah
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		helper.WriteToResponseBody(writer, webResponse)
	}

}
