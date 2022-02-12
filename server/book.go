package server

import (
	"librarySysfo/database"
	"net/http"
)

func allBook(w http.ResponseWriter, r *http.Request) {
	var book []book
	database.DB.Find(&book)
	response := responseParam{
		W:      w,
		Body:   respToByte("success", book, http.StatusAccepted),
		Status: http.StatusAccepted,
	}
	responseFormatter(response, "")
}
