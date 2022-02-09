package seeds

import (
	"crypto/md5"
	"fmt"
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
			Password:  fmt.Sprintf("%x", md5.Sum([]byte("12312312"))),
			LastLogin: &now,
		})
	}

	database.DB.Create(&data)
	util.InfoPrint(2, task)
}
