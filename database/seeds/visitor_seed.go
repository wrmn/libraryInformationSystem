package seeds

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func visitorSeed() {
	task := "Seeding Table Visitor"
	t := time.Now()
	data := []models.Visitor{}
	rand.Seed(t.UnixNano())
	util.InfoPrint(1, task)
	for c := 0; c < 200; c++ {
		singleData := models.Visitor{
			Id:      c + 1,
			LoginAt: t,
			Method:  rand.Intn(4),
			Purpose: gofakeit.Sentence(2),
		}
		if c < 120 {
			uid := rand.Intn(44) + 6
			singleData.UserId = &uid
		} else {
			uid := c - 120 + 1
			singleData.GuestId = &uid
		}
		data = append(data, singleData)
	}

	database.DB.Create(&data)

	util.InfoPrint(2, task)
}
