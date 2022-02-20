package server

import (
	"fmt"
	"librarySysfo/database"
	"librarySysfo/util"
	"net/http"

	"github.com/gorilla/mux"
)

func allBook(w http.ResponseWriter, r *http.Request) {
	util.InfoPrint(3, "New request for book")
	var book []book
	title := r.FormValue("title")
	author := r.FormValue("author")
	var query, titleSearch, authorSearch string

	if title != "" {
		util.InfoPrint(3, "title search included")
		titleSearch = fmt.Sprintf("%%%s%%", title)
		query = query + "LOWER(title) LIKE LOWER(?)"
	}
	if author != "" {
		util.InfoPrint(3, "author search included")
		authorSearch = fmt.Sprintf("%%%s%%", author)
		if query != "" {
			query = query + " AND "
		}
		query = query + "LOWER(author) LIKE LOWER(?)"
	}

	searchDb := database.DB

	if title != "" && author != "" {
		searchDb = searchDb.Where(query, titleSearch, authorSearch)
	} else if title != "" {
		searchDb = searchDb.Where(query, titleSearch)
	} else if author != "" {
		searchDb = searchDb.Where(query, authorSearch)
	}

	searchDb.Find(&book)
	util.InfoPrint(3, "returning response")
	response := responseParam{
		W:      w,
		Body:   respToByte("success", book, http.StatusOK),
		Status: http.StatusOK,
	}
	responseFormatter(response, "")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var book book
	util.InfoPrint(3, fmt.Sprintf("looking for book with id %s", id))
	database.DB.First(&book, id)
	util.InfoPrint(3, "returning response")
	response := responseParam{
		W:      w,
		Body:   respToByte("success", book, http.StatusAccepted),
		Status: http.StatusAccepted,
	}
	responseFormatter(response, "")
}
