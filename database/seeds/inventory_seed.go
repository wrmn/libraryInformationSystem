package seeds

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func inventorySeed() {
	var regSerial, invSerial int64

	task := "Seeding Table Inventory"
	util.InfoPrint(1, task)
	rand.Seed(time.Now().UnixNano())
	for c := 0; c < 30; c++ {
		regId := rand.Intn(6) + 1
		database.DB.
			Model(&models.Inventory{}).
			Where("registration_id=?", regId).
			Count(&invSerial)
		database.DB.
			Model(&models.Book{}).
			Where("registration_id", regId).
			Count(&regSerial)
		data := models.Inventory{
			Id:             c + 1,
			RegistrationId: regId,
			SerialNumber:   int(regSerial+invSerial) + 1,
			Name:           gofakeit.Word(),
			Category:       rand.Intn(4) + 1,
			Status:         rand.Intn(5) + 1,
			Description:    gofakeit.Sentence(12),
		}
		database.DB.Create(&data)
	}

	util.InfoPrint(2, task)

}
