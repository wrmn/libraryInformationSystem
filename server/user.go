package server

import (
	"encoding/json"
	"fmt"
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	util.InfoPrint(1, fmt.Sprintf("New Request %s", r.URL.Path))
	var creds models.User
	var data user
	var count int64

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		badRequest(w)
		return
	}

	database.DB.
		Find(&creds, "username = ? AND password = ?", data.Username, data.Password).
		Count(&count)

	if count <= 0 {
		util.InfoPrint(5, "user not found")
		response := responseParam{
			W:      w,
			Body:   respToByte("error", "Record not found", http.StatusUnauthorized),
			Status: http.StatusUnauthorized,
		}
		responseFormatter(response, " ")
		return
	}

	util.InfoPrint(3, fmt.Sprintf("user Authorized, username %s", data.Username))
	util.InfoPrint(1, "Creating token")

	result := tokenCred{}
	result.AccessToken, err = createToken(creds)
	if err != nil {
		log.Fatal(err)
		util.InfoPrint(5, err.Error())
		return
	}

	util.InfoPrint(2, "Creating token")
	util.InfoPrint(1, "Creating refresh token")

	result.RefreshToken, err = createRefreshToken(result)
	if err != nil {
		log.Fatal(err)
		util.InfoPrint(5, err.Error())
		return
	}
	util.InfoPrint(2, "Creating refresh token")

	response := responseParam{
		W:      w,
		Body:   respToByte("success", result, http.StatusAccepted),
		Status: http.StatusAccepted,
	}
	responseFormatter(response, "")
	util.InfoPrint(2, fmt.Sprintf("New Request %s", r.URL.Path))

}

func dashboard(w http.ResponseWriter, r *http.Request) {
	util.InfoPrint(3, fmt.Sprintf("New Request %s", r.URL.Path))
	token, err := inspectToken(w, r)
	if err {
		return
	}
	user := token.Claims.(*credential)
	json.NewEncoder(w).Encode(fmt.Sprintf("%x Dashboard", user))
}
