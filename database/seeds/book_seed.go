package seeds

import (
	"fmt"
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func bookSeed() {
	var regSerial, ddcSerial, invSerial int64

	task := "Seeding Table Book"
	rand.Seed(time.Now().UnixNano())

	util.InfoPrint(1, task)

	for c := 0; c < 2000; c++ {
		regId := rand.Intn(6) + 1
		ddc := util.IntToDdc(rand.Intn(1000))
		database.DB.
			Model(&models.Inventory{}).
			Where("registration_id=?", regId).
			Count(&invSerial)
		database.DB.
			Model(&models.Book{}).
			Where("registration_id", regId).
			Count(&regSerial)
		database.DB.
			Model(&models.Book{}).
			Where("ddc_no=?", ddc).
			Count(&ddcSerial)

		data := models.Book{
			Id:             c + 1,
			RegistrationId: regId,
			SerialNumber:   int(invSerial+regSerial) + 1,
			DdcNo:          ddc,
			DdcOrder:       int(ddcSerial) + 1,
			Title:          gofakeit.Sentence(4),
			Author:         fmt.Sprintf("%s %s", gofakeit.FirstName(), gofakeit.Word()),
			Publisher:      fmt.Sprintf("%s %s", gofakeit.LastName(), gofakeit.Word()),
			Availability:   true,
			Status:         rand.Intn(5) + 1,
			Price:          (rand.Intn(250) + 50) * 1000,
			CoverFile:      "default.jpg",
		}
		database.DB.Create(&data)
	}

	util.InfoPrint(2, task)
}
