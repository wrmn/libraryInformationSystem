package server

import (
	"encoding/json"
	"fmt"
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
	"net/http"
	"time"
)

func checkinMember(w http.ResponseWriter, r *http.Request) {
	var data models.Visitor
	var dataCount int64
	util.InfoPrint(3, "new checkin for member")
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		util.InfoPrint(5, err.Error())
		badRequest(w)
		return
	}

	visitor := database.DB.Model(&models.Visitor{})
	visitor.Count(&dataCount)
	data.Id = int(dataCount + 1)
	data.LoginAt = time.Now()

	if err := database.DB.Create(&data).Error; err != nil {
		fmt.Println(err.Error())
	}
}

func checkinGuest(w http.ResponseWriter, r *http.Request) {
	var data models.Guest
	var dataCount, visitCount int64

	util.InfoPrint(3, "new checkin for guest")
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		util.InfoPrint(5, err.Error())
		badRequest(w)
		return
	}

	visitor := database.DB.Model(&models.Visitor{})
	guest := database.DB.Model(&models.Guest{})
	visitor.Count(&visitCount)
	guest.Count(&dataCount)

	data.Id = int(dataCount + 1)
	data.Visitor.Id = int(visitCount + 1)
	data.Visitor.GuestId = &data.Id
	data.Visitor.LoginAt = time.Now()

	if err := database.DB.Create(&data).Error; err != nil {
		fmt.Println(err.Error())
	}
}
