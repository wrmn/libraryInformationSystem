package server

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
	"log"
	"net/http"
	"regexp"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	util.InfoPrint(1, fmt.Sprintf("New Request %s", r.URL.Path))
	var creds models.User
	var data userIn

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		badRequest(w)
		return
	}

	hashPasswd := md5.Sum([]byte(data.Password))
	passwd := fmt.Sprintf("%x", hashPasswd)
	database.DB.
		Where("username = ? AND password = ?", data.Username, passwd).
		Find(&creds)
	database.DB.Where("id = ?", creds.Id).Find(&creds.Employee)

	if data.Username != creds.Username {
		util.InfoPrint(5, fmt.Sprintf("user %s not found", data.Username))
		response := responseParam{
			W:      w,
			Body:   respToByte("error", "Record not found", http.StatusUnauthorized),
			Status: http.StatusUnauthorized,
		}
		responseFormatter(response, " ")
		return
	}

	util.InfoPrint(3, fmt.Sprintf("User Authorized, username %s", data.Username))
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
	user := token.Claims.(*credential).Username
	json.NewEncoder(w).Encode(fmt.Sprintf("%s Dashboard", user))
}

func register(w http.ResponseWriter, r *http.Request) {
	util.InfoPrint(1, fmt.Sprintf("New Request %s", r.URL.Path))
	var data models.User
	timeNow := time.Now()
	util.InfoPrint(1, "Reading request body")
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		util.InfoPrint(5, err.Error())
		badRequest(w)
		return
	}

	hashPassword := md5.Sum([]byte(data.Password))
	data.Password = fmt.Sprintf("%x", hashPassword)
	data.LastLogin = &timeNow

	var dataCount int64
	user := database.DB.Model(&models.User{})
	user.Count(&dataCount)
	data.Id = int(dataCount + 1)
	if err := database.DB.Create(&data).Error; err != nil {
		isUname, _ := regexp.MatchString(".*duplicate.*username.*", err.Error())
		isEmail, _ := regexp.MatchString(".*duplicate.*email.*", err.Error())
		if isUname {
			response := responseParam{
				W:      w,
				Body:   respToByte("error", "Username has been used by another account", http.StatusBadRequest),
				Status: http.StatusBadRequest,
			}
			responseFormatter(response, "")
			return
		} else if isEmail {
			response := responseParam{
				W:      w,
				Body:   respToByte("error", "Email has been used by another account", http.StatusBadRequest),
				Status: http.StatusBadRequest,
			}
			responseFormatter(response, "")
			return
		} else {
			response := responseParam{
				W:      w,
				Body:   respToByte("error", "Can't create account right now", http.StatusBadRequest),
				Status: http.StatusBadRequest,
			}
			responseFormatter(response, "")
			return
		}
	} else {
		msg := fmt.Sprintf("User %s has been created", data.Username)
		response := responseParam{
			W:      w,
			Body:   respToByte("success", msg, http.StatusOK),
			Status: http.StatusOK,
		}
		responseFormatter(response, "")
		return
	}
}

func info(w http.ResponseWriter, r *http.Request) {
	util.InfoPrint(3, fmt.Sprintf("New Request %s", r.URL.Path))
	token, err := inspectToken(w, r)
	if err {
		return
	}
	user := token.Claims.(*credential)

	response := responseParam{
		W:      w,
		Body:   respToByte("success", user, http.StatusAccepted),
		Status: http.StatusAccepted,
	}
	responseFormatter(response, "")
}
