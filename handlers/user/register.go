package user

import (
	"encoding/json"
	"log"
	"net/http"
	"url_short/database"
	"url_short/util"
)

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var usr database.User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

	if usr.CheckUser() {
		http.Error(w, "User Already Exists", http.StatusForbidden)
		return
	}

	err = usr.CreateUser()

	if err != nil {
		http.Error(w, "Something Went Wrong! Try again later!", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	util.SendData(w, http.StatusCreated, "Account Created Successfully. Please Login")

}
