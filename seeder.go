package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/schollz/progressbar/v3"
)

//make query for insert data assetRecording
func (i InsertParam) AssetRecordingSeeder() {

}

//make query for insert data user
func (i InsertParam) DdcSeeder() {
	var num string
	bar := progressbar.Default(1000)
	ddc := Ddc{}
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
		if c < 10 {
			num = fmt.Sprintf("00%d", c)
		} else if c < 100 {
			num = fmt.Sprintf("0%d", c)
		} else {
			num = strconv.Itoa(c)
		}

		ddc.Ddc = num
		ddc.Group = category[(c / 100)]
		ddc.Name = fmt.Sprintf("%s %s", faker.Word(), faker.Word())

		query := composeInsert(i.TableName, ddc)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, "")
		}
		bar.Add(1)
	}
}

//make query for insert data user
func (i InsertParam) UserSeeder() {
	user := User{}
	bar := progressbar.Default(6)
	t := time.Now().Format("2006-01-02 15:04:05")
	for c := 0; c < 6; c++ {
		user.Id = c + 1
		user.Username = faker.Username()
		user.Password = fmt.Sprintf("%x", md5.Sum([]byte(faker.Word())))
		user.CreatedAt, user.LastLogin, user.UpdatedAt = t, t, t
		query := composeInsert(i.TableName, user)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, "")
		}
		bar.Add(1)
	}
}
