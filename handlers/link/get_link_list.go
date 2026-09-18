package link

import (
	"log"
	"net/http"
	"strings"
	"url_short/database"
	"url_short/util"
)

func (h *Handler) GetLinkList(w http.ResponseWriter, r *http.Request) {
	var linkStruct database.Link

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

	linkStruct.UserId = *userId

	linkList := linkStruct.GetList()

	util.SendData(w, http.StatusOK, linkList)

}
