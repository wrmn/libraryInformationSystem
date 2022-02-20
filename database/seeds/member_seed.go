package seeds

import (
	"fmt"
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func memberSeed() {
	data := []models.Member{}
	task := "Seeding Table Member"
	util.InfoPrint(1, task)
	now := time.Now()
	rand.Seed(now.UnixNano())
	for c := 0; c < 44; c++ {
		var identityNo string
		birthDate := util.DateRandom("1900-01-01", "2016-01-01")
		address := gofakeit.Address()
		address2 := fmt.Sprintf("%s, %s",
			address.City,
			address.Country,
		)
		identityType := rand.Intn(4) + 1
		if identityType == 3 {
			identityNo = util.RandDigit(12)
		} else if identityType == 4 {
			identityNo = fmt.Sprintf("%s/%s/%s",
				util.RandDigit(5),
				strings.ToUpper(gofakeit.LetterN(2)),
				birthDate.Format("2006"),
			)
		} else {
			identityNo = fmt.Sprintf("%s%s%s",
				util.RandDigit(6),
				birthDate.Format("02012006"),
				util.RandDigit(4),
			)
		}
		verified := (rand.Intn(2) == 0)

		singleData := models.Member{
			Id:            c + 7,
			Name:          gofakeit.Name(),
			Gender:        util.Gender[rand.Intn(2)],
			PlaceOfBirth:  gofakeit.Address().City,
			DateOfBirth:   birthDate.Format(util.Dmy),
			Address1:      address.Street,
			Address2:      &address2,
			Profession:    rand.Intn(8) + 1,
			Institution:   gofakeit.Company(),
			PhoneNo:       gofakeit.Phone(),
			IsWhatsapp:    (rand.Intn(2) == 0),
			IdentityType:  identityType,
			IdentityNo:    identityNo,
			IdentityFile:  "default.jpg",
			PhotoFile:     "default.jpg",
			AgreementFile: "default.jpg",
		}

		if verified {
			singleData.VerifiedAt = &now
		}

		data = append(data, singleData)
	}
	database.DB.Create(&data)
	util.InfoPrint(2, task)
}
