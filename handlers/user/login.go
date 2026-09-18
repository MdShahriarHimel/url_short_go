package user

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"url_short/database"
	"url_short/util"
)

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var usr database.User

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	verifiedUser, err := usr.VerifyUser()

	if err != nil {
		log.Println(err)
		http.Error(w, "Oops! something went wrong!", http.StatusInternalServerError)
		return
	}

	if verifiedUser == nil {
		http.Error(w, "Invalid Credentials", http.StatusNotFound)
		return
	}

	payload := util.Payload{
		SUB:      verifiedUser.ID,
		UserName: verifiedUser.UserName,
		Email:    verifiedUser.Email,
	}

	jwt_token, err := util.CreateJwt(payload)

	if err != nil {
		http.Error(w, "Oops! Something Went Wrong", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// refresh token

	refreshToken, err := util.CreateRefreshToken()
	if err != nil {
		http.Error(w, "Oops! Something Went Wrong", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	//Refresh token
	refreshTokenStruct := database.RefreshToken{
		UserId:    verifiedUser.ID,
		TokenHash: refreshToken,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(3 * time.Minute).Unix(),
		Revoked:   false,
	}

	refreshTokenStruct.Save()

	// fmt.Println(refreshTokenStruct)

	data := map[string]any{
		"jwt_token":     *jwt_token,
		"refresh_token": refreshToken,
	}

	util.SendData(w, http.StatusOK, data)

}
