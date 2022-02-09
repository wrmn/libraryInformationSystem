package seeds

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func guestSeed() {
	task := "Seeding Table Guest"
	util.InfoPrint(1, task)
	data := []models.Guest{}
	rand.Seed(time.Now().UnixNano())

	for c := 0; c < 80; c++ {
		data = append(data, models.Guest{
			Id:          c + 1,
			Name:        gofakeit.Name(),
			Gender:      util.Gender[rand.Intn(2)],
			Address:     gofakeit.Address().Address,
			Profession:  rand.Intn(8) + 1,
			Institution: gofakeit.Company(),
		})
	}
	database.DB.Create(&data)
	util.InfoPrint(2, task)
}
