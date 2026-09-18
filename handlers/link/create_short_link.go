package link

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"url_short/database"
	"url_short/util"
)

func (h *Handler) CreateShortLink(w http.ResponseWriter, r *http.Request) {
	var linkStruct database.Link

	err := json.NewDecoder(r.Body).Decode(&linkStruct)
	if err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
	}

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

	link, stored := linkStruct.StoreShortLink()
	if stored == false {
		util.SendData(w, http.StatusConflict, *link)
		return
	}

	util.SendData(w, http.StatusOK, linkStruct)

}
