package seeds

import (
	"librarySysfo/database"
	"librarySysfo/database/models"
	"librarySysfo/util"

	"github.com/brianvoe/gofakeit/v6"
)

func ddcSeed() {
	data := []models.Ddc{}
	task := "Seeding Table Ddc"
	util.InfoPrint(1, task)
	category := []string{
		"karya umum",
		"filsafat",
		"agama",
		"ilmu sosial",
		"bahasa",
		"ilmu murni",
		"ilmu terapan",
		"seni dan olahraga",
		"kesusastraan",
		"sejarah dan geografi",
	}

	for c := 0; c < 1000; c++ {
		data = append(data, models.Ddc{
			Ddc:   util.IntToDdc(c),
			Group: category[c/100],
			Name:  gofakeit.Sentence(2),
		})
	}

	database.DB.Create(&data)
	util.InfoPrint(2, task)

}
