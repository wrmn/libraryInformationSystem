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
	util.InfoPrint(3, fmt.Sprintf("New Request %s", r.URL.Path))

	var creds models.User
	var some user
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		badRequest(w)
		return
	}
	// userPassword, ok := userdb[creds.Username]

	// // if user exist, verify the password
	// if !ok || userPassword != creds.Password {
	// 	unAuthorized(w)
	// 	return
	// }
	database.DB.Where(&creds).First(&some)
	fmt.Printf("%+v\n", some)

	// TODO: change auth from database
	result := token{}
	result.AccessToken, err = createToken(creds)
	if err != nil {
		log.Fatal(err)
		util.InfoPrint(5, err.Error())
		return
	}
	result.RefreshToken, err = createRefreshToken(result)
	if err != nil {
		log.Fatal(err)
		util.InfoPrint(5, err.Error())
		return
	}

	msg := respToByte("success", result, http.StatusAccepted)
	response := responseParam{
		W:      w,
		Body:   msg,
		Status: http.StatusAccepted,
	}
	responseFormatter(response)
}

func refresh(w http.ResponseWriter, r *http.Request) {
	var creds token
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		badRequest(w)
		return
	}
	user, err := validateRefreshToken(creds)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid Token")
		return
	}

	creds.AccessToken, err = createToken(user)
	if err != nil {
		json.NewEncoder(w).Encode("Unable to create access token")
		return
	}
	msg := respToByte("Success", creds, http.StatusOK)
	response := responseParam{
		W:      w,
		Body:   msg,
		Status: http.StatusAccepted,
	}
	responseFormatter(response)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	token, err := inspectToken(w, r)
	if err != nil {
		return
	}
	user := token.Claims.(*credential)
	json.NewEncoder(w).Encode(fmt.Sprintf("%s Dashboard", user.Username))
}
