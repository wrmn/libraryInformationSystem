package seeds

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
	"math/rand"
	"time"
)

func borrowSeeder() {
	task := "Seeding Table Borrow"
	rand.Seed(time.Now().UnixNano())
	data := []models.Borrow{}
	now := time.Now().Format(util.Dmy)

	util.InfoPrint(1, task)
	for c := 0; c < 600; c++ {
		fineStatus := (rand.Intn(2) == 0)
		singleData := models.Borrow{
			Id:           c + 1,
			AdminId:      rand.Intn(6) + 1,
			MemberId:     rand.Intn(44) + 6,
			BookId:       rand.Intn(2000) + 1,
			FineStatus:   fineStatus,
			DateOfBorrow: util.DateRandom("2015-01-01", now),
			DateOfReturn: util.DateRandom("2015-01-01", now),
		}

		if fineStatus {
			status := (rand.Intn(2) == 0)
			if status {
				singleData.DateOfPayment = util.DateRandom("2015-01-01", now)
			}
		}
		data = append(data, singleData)
	}

	database.DB.Create(&data)
	util.InfoPrint(2, task)
}
