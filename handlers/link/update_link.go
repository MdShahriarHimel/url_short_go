package link

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"url_short/database"
	"url_short/util"
)

func (h *Handler) UpdateLink(w http.ResponseWriter, r *http.Request) {
	var linkStruct database.Link

	err := json.NewDecoder(r.Body).Decode(&linkStruct)
	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
		return
	}

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

	updated := linkStruct.Update()
	if !updated {
		http.Error(w, "Invalid Link to update", http.StatusBadRequest)
	}

	util.SendData(w, http.StatusOK, "updated Successfully")

}
