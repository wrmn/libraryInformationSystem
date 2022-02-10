package server

import (
	"fmt"
	"librarySysfo/util"
	"net/http"
)

func notFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		util.InfoPrint(3, fmt.Sprintf("New Request %s", r.URL.Path))
		util.InfoPrint(3, "New Request")
		warning := fmt.Sprintf("Request for %s is not found", r.URL.Path)
		util.InfoPrint(4, warning)
		msg := respToByte("error", warning, http.StatusNotFound)
		response := responseParam{
			W:      w,
			Body:   msg,
			Status: http.StatusNotFound,
		}
		responseFormatter(response)
	})
}

func notAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		util.InfoPrint(3, fmt.Sprintf("New Request %s", r.URL.Path))
		warning := fmt.Sprintf("Request for %s is not allowed with method %s", r.URL.Path, r.Method)
		util.InfoPrint(4, warning)
		msg := respToByte("error", warning, http.StatusMethodNotAllowed)
		response := responseParam{
			W:      w,
			Body:   msg,
			Status: http.StatusMethodNotAllowed,
		}
		responseFormatter(response)
	})
}

func unAuthorized(w http.ResponseWriter) {
	info := "No match credential"
	util.InfoPrint(4, info)
	msg := respToByte("error", info, http.StatusUnauthorized)
	response := responseParam{
		W:      w,
		Body:   msg,
		Status: http.StatusUnauthorized,
	}
	responseFormatter(response)
}

func badRequest(w http.ResponseWriter) {
	info := "Bad format request"
	util.InfoPrint(4, info)
	msg := respToByte("error", info+", check documentation!", http.StatusBadRequest)
	response := responseParam{
		W:      w,
		Body:   msg,
		Status: http.StatusBadRequest,
	}
	responseFormatter(response)
}
