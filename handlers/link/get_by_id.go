package link

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"url_short/database"
	"url_short/util"
)

func (h *Handler) GetLinkById(w http.ResponseWriter, r *http.Request) {
	var linkStruct database.Link

	id, err := strconv.Atoi(r.PathValue("link_id"))
	if err != nil {
		http.Error(w, "Invalid parameter or link id", http.StatusBadRequest)
		log.Println(err)
		return
	}

	linkStruct.ID = id

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

	link := linkStruct.GetById()

	if link == nil {
		http.Error(w, "No Link Found!", http.StatusNotFound)
		return
	}

	util.SendData(w, http.StatusOK, *link)

}
