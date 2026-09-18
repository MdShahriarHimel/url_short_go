package link

import (
	"log"
	"net/http"
	"strings"
	"url_short/database"
	"url_short/util"
)

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	var link database.Link
	shortCode := r.PathValue("short_code")

	if len(shortCode) == 0 {
		http.Error(w, "Invalid short Url", http.StatusBadRequest)
		return
	}

	link.ShortCode = shortCode

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

	link.UserId = *userId

	longUrl, found := link.GetLongUrl()
	if !found {
		http.Error(w, "Short Url not valid", http.StatusBadRequest)
		return
	}
	// log.Println(longUrl)
	http.Redirect(w, r, longUrl, http.StatusFound)

}
