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

func employeeSeed() {
	data := []models.Employee{}
	task := "Seeding Table Employee"
	util.InfoPrint(1, task)

	for c := 0; c < 6; c++ {
		birthDate := util.DateRandom("1960-01-01", "200-01-01")
		genderIndex := rand.Intn(2)
		address := gofakeit.Address()
		address2 := fmt.Sprintf("%s, %s",
			address.City,
			address.Country,
		)

		data = append(data, models.Employee{
			Id: c + 1,
			EmployeeNumber: fmt.Sprintf("%s%s%d%s",
				birthDate.Format("20060102"),
				util.DateRandom("1980-01-01", time.Now().Format(util.Dmy)).
					Format("200601"),
				genderIndex,
				util.RandDigit(3),
			),
			Name:         gofakeit.Name(),
			Gender:       util.Gender[genderIndex],
			DateOfBirth:  birthDate,
			PlaceOfBirth: gofakeit.City(),
			Address1:     address.Street,
			Address2:     &address2,
			Division:     rand.Intn(5) + 1,
			Position:     gofakeit.JobTitle(),
		})
	}

	database.DB.Create(&data)
	util.InfoPrint(2, task)
}
