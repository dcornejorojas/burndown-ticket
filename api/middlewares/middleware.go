package middlewares

import (
	"net/http"
	"ticket/api/auth"
	"ticket/api/models"
	"ticket/api/utils"
)

func SetAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.ValidToken(r)
		if err != nil {
			errObj := models.Error{}
			errObj.HasError(true, http.StatusUnauthorized, "Unauthorized")
			utils.ResponseJSON(w, http.StatusUnauthorized, "Token invalido", []string{}, errObj)
			return
		}
		next(w, r)
	}
}
