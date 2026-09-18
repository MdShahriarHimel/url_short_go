package user

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"url_short/config"
	"url_short/database"
	"url_short/util"
)

func (h *Handler) GetNewAccessToken(w http.ResponseWriter, r *http.Request) {
	// check refresh token is valid?

	var refreshTokenStruct database.RefreshToken

	err := json.NewDecoder(r.Body).Decode(&refreshTokenStruct)
	if err != nil {
		http.Error(w, "Not valid json!", http.StatusBadRequest)
		return
	}

	if !refreshTokenStruct.IsValid() {
		http.Error(w, "Not valid refreshToken!", http.StatusBadRequest)
		return
	}

	// revoke refresh token
	refreshTokenStruct.Revoked = false

	// fmt.Println(refreshTokenStruct)

	// regenerate refresh token

	newRefreshToken, err := util.CreateRefreshToken()
	if err != nil {
		http.Error(w, "Something Went Wrong!", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	newRefreshTokenStruct := database.RefreshToken{
		UserId:    refreshTokenStruct.UserId,
		TokenHash: newRefreshToken,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Duration(config.GetConfig().JWT_REFRESH_TOKEN_LIFE_TIME) * time.Minute).Unix(),
		Revoked:   false,
	}

	newRefreshTokenStruct.Save()

	// new Access token

	var user database.User
	user.ID = newRefreshTokenStruct.UserId

	found := user.GetUserById()
	if !found {
		http.Error(w, "Something Went Wrong userid!", http.StatusInternalServerError)
		return
	}

	payload := util.Payload{
		SUB:      user.ID,
		UserName: user.UserName,
		Email:    user.Email,
	}

	jwtToken, err := util.CreateJwt(payload)

	if err != nil {
		http.Error(w, "Something Went Wrong!", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	data := map[string]any{
		"jwt_token":         jwtToken,
		"new_refresh_token": newRefreshToken,
	}

	util.SendData(w, http.StatusOK, data)

}
