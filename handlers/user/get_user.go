package user

import (
	"net/http"
	"url_short/database"
	"url_short/util"
)

func (h *Handler) GetUserList(w http.ResponseWriter, r *http.Request) {
	users := database.GetUserList()

	util.SendData(w, http.StatusOK, users)
}
