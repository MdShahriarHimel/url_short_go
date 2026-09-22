package link

import (
	"net/http"
	"url_short/database"
)

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	var link database.Link
	shortCode := r.PathValue("short_code")

	if len(shortCode) == 0 {
		http.Error(w, "Invalid short Url", http.StatusBadRequest)
		return
	}

	link.ShortCode = shortCode

	longUrl, found := link.GetLongUrl()
	if !found {
		http.Error(w, "Short Url not valid", http.StatusBadRequest)
		return
	}
	// log.Println(longUrl)
	http.Redirect(w, r, longUrl, http.StatusFound)

}
