package link

import (
	"net/http"
	"strconv"
	"url_short/database"
	"url_short/util"
)

func (h *Handler) DeleteLink(w http.ResponseWriter, r *http.Request) {
	var link database.Link
	idStr := r.PathValue("link_id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Link Id", http.StatusBadRequest)
		return
	}

	link.ID = id

	deleted := link.Delete()
	if !deleted {
		http.Error(w, "Link not found!", http.StatusBadRequest)
		return
	}

	util.SendData(w, http.StatusOK, "Link Deleted Successfully!")

}
