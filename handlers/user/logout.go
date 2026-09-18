package user

import (
	"log"
	"net/http"
	"strings"
	"url_short/database"
	"url_short/util"
)

func (h *Handler) LogOut(w http.ResponseWriter, r *http.Request) {
	jwt_token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if len(jwt_token) == 0 {
		http.Error(w, "Invalid Auth Header", http.StatusBadRequest)
		log.Println("Invalid Auth Header at CreateShortLink")
		return
	}

	userId, err := util.GetUserIdFromJwt(jwt_token)
	if err != nil {
		http.Error(w, "Oops! something went wrong!", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	refreshTokenStruct := database.RefreshToken{
		UserId: *userId,
	}

	refreshTokenStruct.RevokeRefreshToken()
}
