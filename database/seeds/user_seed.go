package seeds

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func userSeed() {
	data := []models.User{}
	task := "Seeding Table User"
	util.InfoPrint(1, task)
	now := time.Now()
	rand.Seed(now.UnixNano())
	for c := 0; c < 50; c++ {
		data = append(data, models.User{
			Id:        c + 1,
			Username:  gofakeit.Gamertag(),
			Email:     gofakeit.Email(),
			Password:  "f5bb0c8de146c67b44babbf4e6584cc0",
			LastLogin: &now,
		})
	}

	database.DB.Create(&data)
	util.InfoPrint(2, task)
}
