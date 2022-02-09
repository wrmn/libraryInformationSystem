package seeds

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"

	"github.com/brianvoe/gofakeit/v6"
)

func assetRecordSeed() {
	data := []models.AssetRecord{}
	task := "Seeding Table assetRecord"
	rand.Seed(time.Now().UnixNano())

	util.InfoPrint(1, task)
	for c := 0; c < 6; c++ {
		data = append(data, models.AssetRecord{
			Id:                 c + 1,
			AdminId:            rand.Intn(6) + 1,
			RegistrationNumber: strings.ToUpper(fmt.Sprintf("%s/%s/%d", gofakeit.Letter(), gofakeit.LetterN(3), gofakeit.Year())),
			RegistrationDate:   util.DateRandom("2015-01-01", time.Now().Format(util.Dmy)),
			Source:             gofakeit.Word(),
		})
	}

	database.DB.Create(&data)
	util.InfoPrint(2, task)
}
