package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/schollz/progressbar/v3"
)

//make query for insert data assetRecord
func (i QueryParam) AssetRecordSeeder() {
	data := AssetRecord{}
	bar := progressbar.Default(6)
	t := time.Now().Format("2006-01-02 15:04:05")
	for c := 0; c < 6; c++ {
		data.Id = c + 1
		data.AdminId = rand.Intn(6) + 1
		data.RegistrationNumber = fmt.Sprintf("%s/%s/%s", faker.Century(), faker.Currency(), faker.YearString())
		data.RegistrationDate = faker.Date()
		data.Source = faker.Word()
		data.CreatedAt, data.UpdatedAt = t, t
		query := composeInsert(i.TableName, data)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, query)
		}
		bar.Add(1)
	}
}

func (i QueryParam) BookSeeder() {
	data := Book{}
	bar := progressbar.Default(2000)
	t := time.Now().Format("2006-01-02 15:04:05")
	rand.Seed(time.Now().UnixNano())
	countSerial := SearchParam{
		Column: "serialNumber",
	}
	countDdc := SearchParam{
		Column: "ddc",
	}
	for c := 0; c < 2000; c++ {
		data.Id = c + 1
		data.RegistrationId = rand.Intn(6) + 1
		countSerial.Value = data.RegistrationId
		data.SerialNumber = selectCount(i, countSerial) + 1
		data.Ddc = intToDdc(rand.Intn(1000))
		countDdc.Value = data.Ddc
		data.DdcOrder = selectCount(i, countDdc) + 1
		data.Title = fmt.Sprintf("%s %s %s", faker.Word(), faker.Word(), faker.Word())
		data.Author = fmt.Sprintf("%s %s", faker.FirstName(), faker.Word())
		data.Publisher = fmt.Sprintf("%s %s", faker.LastName(), faker.Word())
		data.Availability = true
		data.Price = (rand.Intn(250) + 50) * 1000
		data.CreatedAt, data.UpdatedAt = t, t
		query := composeInsert(i.TableName, data)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, "")
		}
		bar.Add(1)
	}
}
func (i QueryParam) BorrowSeeder() {}

//make query for insert data user
func (i QueryParam) DdcSeeder() {
	data := Ddc{}
	bar := progressbar.Default(1000)
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
		data.Ddc = intToDdc(c)
		data.Group = category[(c / 100)]
		data.Name = fmt.Sprintf("%s %s", faker.Word(), faker.Word())

		query := composeInsert(i.TableName, data)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, "")
		}
		bar.Add(1)
	}
}

func (i QueryParam) EmployeeSeeder()  {}
func (i QueryParam) GuestSeeder()     {}
func (i QueryParam) InventorySeeder() {}
func (i QueryParam) MemberSeeder()    {}

//make query for insert data user
func (i QueryParam) UserSeeder() {
	data := User{}
	bar := progressbar.Default(6)
	t := time.Now().Format("2006-01-02 15:04:05")
	for c := 0; c < 6; c++ {
		data.Id = c + 1
		data.Username = faker.Username()
		data.Password = fmt.Sprintf("%x", md5.Sum([]byte(faker.Word())))
		data.CreatedAt, data.LastLogin, data.UpdatedAt = t, t, t
		query := composeInsert(i.TableName, data)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, "")
		}
		bar.Add(1)
	}
}

func (i QueryParam) VisitorSeeder() {}
