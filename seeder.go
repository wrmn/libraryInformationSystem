package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/schollz/progressbar/v3"
)

//make query for insert data assetRecord
func (i QueryParam) AssetRecordSeeder() {
	data := AssetRecord{}
	bar := progressbar.Default(6)
	t := time.Now().Format(dmyhms)
	rand.Seed(time.Now().UnixNano())
	for c := 0; c < 6; c++ {
		data.Id = c + 1
		data.AdminId = rand.Intn(6) + 1
		data.RegistrationNumber = strings.ToUpper(fmt.Sprintf("%s/%s/%d", gofakeit.Letter(), gofakeit.LetterN(3), gofakeit.Year()))
		data.RegistrationDate = dateRandom("2015-01-01", time.Now().Format(dmy)).Format(dmy)
		data.Source = gofakeit.Word()
		data.CreatedAt, data.UpdatedAt = t, t
		query := composeInsert(i.TableName, data)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, "Error while executing : "+query)
		}
		bar.Add(1)
	}
}

func (i QueryParam) BookSeeder() {
	data := Book{}
	bar := progressbar.Default(2000)
	t := time.Now().Format(dmyhms)
	rand.Seed(time.Now().UnixNano())
	countSerial := SearchParam{
		Column: "registrationId",
	}
	countDdc := SearchParam{
		Column: "ddc",
	}
	for c := 0; c < 2000; c++ {
		data.Id = c + 1
		data.RegistrationId = rand.Intn(6) + 1
		countSerial.Value = data.RegistrationId
		data.SerialNumber = totalSerial(i.Db, countSerial) + 1
		data.Ddc = intToDdc(rand.Intn(1000))
		countDdc.Value = data.Ddc
		data.DdcOrder = selectCount(i, countDdc) + 1
		data.Title = fmt.Sprintf("%s %s %s", gofakeit.Word(), gofakeit.Word(), gofakeit.Word())
		data.Author = fmt.Sprintf("%s %s", gofakeit.FirstName(), gofakeit.Word())
		data.Publisher = fmt.Sprintf("%s %s", gofakeit.LastName(), gofakeit.Word())
		data.Availability = true
		data.Price = (rand.Intn(250) + 50) * 1000
		data.CreatedAt, data.UpdatedAt = t, t
		query := composeInsert(i.TableName, data)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, "Error while executing : "+query)
		}
		bar.Add(1)
	}
}

func (i QueryParam) BorrowSeeder() {
	data := Borrow{}
	bar := progressbar.Default(600)
	rand.Seed(time.Now().UnixNano())
	now := time.Now().Format(dmy)
	for c := 0; c < 600; c++ {
		data.Id = c + 1
		data.AdminId = rand.Intn(6) + 1
		data.MemberId = rand.Intn(44) + 6
		data.BookId = rand.Intn(2000) + 1
		data.FineStatus = (rand.Intn(4)%2 == 0)
		data.DateOfBorrow = dateRandom("2015-01-01", now).Format(dmyhms)
		data.DateOfReturn = dateRandom("2015-01-01", now).Format(dmyhms)
		data.DateOfPayment = dateRandom("2015-01-01", now).Format(dmyhms)
		query := composeInsert(i.TableName, data)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, "Error while executing : "+query)
		}
		bar.Add(1)
	}

}

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
		data.Name = fmt.Sprintf("%s %s", gofakeit.Word(), gofakeit.Word())
		query := composeInsert(i.TableName, data)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, "Error while executing : "+query)
		}
		bar.Add(1)
	}
}

func (i QueryParam) EmployeeSeeder() {
	data := Employee{}
	bar := progressbar.Default(6)
	rand.Seed(time.Now().UnixNano())

	for c := 0; c < 6; c++ {
		birthDate := dateRandom("1960-01-01", "2000-01-01")
		address := gofakeit.Address()

		genderIndex := rand.Intn(2)
		data.Id = c + 1
		data.EmployeeNumber = fmt.Sprintf("%s%s%d%s",
			birthDate.Format("20060102"),
			dateRandom("1980-01-01", time.Now().Format(dmy)).Format("200601"),
			genderIndex,
			randDigit(3),
		)
		data.Name = gofakeit.Name()
		data.Gender = gender[genderIndex]
		data.PlaceOfBirth = gofakeit.Address().City
		data.DateOfBirth = birthDate.Format(dmy)
		data.Address1 = address.Street
		data.Address2 = fmt.Sprintf("%s, %s",
			address.City,
			address.Country,
		)
		data.Division = strconv.Itoa(rand.Intn(5) + 1)
		data.Position = gofakeit.Job().Title
		query := composeInsert(i.TableName, data)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, "Error while executing : "+query)
		}
		bar.Add(1)
	}

}

func (i QueryParam) GuestSeeder() {
	data := Guest{}
	bar := progressbar.Default(80)
	rand.Seed(time.Now().UnixNano())

	for c := 0; c < 80; c++ {
		data.Id = c + 1
		data.Name = gofakeit.Name()
		data.Gender = gender[rand.Intn(2)]
		data.Address = gofakeit.Address().Address
		data.Profession = strconv.Itoa(rand.Intn(8) + 1)
		data.Institution = gofakeit.Company()
		query := composeInsert(i.TableName, data)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, "Error while executing : "+query)
		}
		bar.Add(1)
	}
}

func (i QueryParam) InventorySeeder() {
	data := Inventory{}
	bar := progressbar.Default(30)
	rand.Seed(time.Now().UnixNano())
	countSerial := SearchParam{
		Column: "registrationId",
	}

	for c := 0; c < 30; c++ {
		data.Id = c + 1
		data.RegistrationId = rand.Intn(6) + 1
		countSerial.Value = data.RegistrationId
		data.SerialNumber = totalSerial(i.Db, countSerial) + 1
		data.Name = gofakeit.Word()
		data.Category = strconv.Itoa(rand.Intn(4) + 1)
		data.Status = strconv.Itoa(rand.Intn(5) + 1)
		data.Description = gofakeit.Sentence(12)
		query := composeInsert(i.TableName, data)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, "Error while executing : "+query)
		}
		bar.Add(1)
	}

}

func (i QueryParam) MemberSeeder() {
	data := Member{}
	bar := progressbar.Default(44)
	t := time.Now().Format(dmyhms)
	rand.Seed(time.Now().UnixNano())
	for c := 0; c < 44; c++ {
		birthDate := dateRandom("1900-01-01", "2016-01-01")
		address := gofakeit.Address()

		data.Id = c + 7
		data.Name = gofakeit.Name()
		data.Gender = gender[rand.Intn(2)]
		data.PlaceOfBirth = gofakeit.Address().City
		data.DateOfBirth = birthDate.Format(dmy)
		data.Address1 = address.Street
		data.Address2 = fmt.Sprintf("%s, %s",
			address.City,
			address.Country,
		)
		data.Profession = strconv.Itoa(rand.Intn(8) + 1)
		data.Institution = gofakeit.Company()
		data.PhoneNo = gofakeit.Phone()
		data.IsWhatsapp = (rand.Intn(4)%2 == 0)
		data.IdentityType = strconv.Itoa(rand.Intn(4) + 1)
		if data.IdentityType == "3" {
			data.IdentityNo = randDigit(12)
		} else if data.IdentityType == "4" {
			data.IdentityNo = fmt.Sprintf("%s/%s/%s",
				randDigit(5),
				strings.ToUpper(gofakeit.LetterN(2)),
				birthDate.Format("2006"),
			)
		} else {
			data.IdentityNo = fmt.Sprintf("%s%s%s",
				randDigit(6),
				birthDate.Format("02012006"),
				randDigit(4),
			)
		}
		data.IdentityFile = "default.jpg"
		data.PhotoFile = "default.jpg"
		data.AgreementFile = "default.pdf"
		data.CreatedAt, data.VerifiedAt, data.UpdatedAt = t, t, t
		query := composeInsert(i.TableName, data)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, "Error while executing : "+query)
		}
		bar.Add(1)
	}

}

//make query for insert data user
func (i QueryParam) UserSeeder() {
	data := User{}
	bar := progressbar.Default(50)
	t := time.Now().Format(dmyhms)
	for c := 0; c < 50; c++ {
		data.Id = c + 1
		data.Username = gofakeit.Gamertag()
		data.Password = fmt.Sprintf("%x", md5.Sum([]byte(gofakeit.BuzzWord())))
		data.CreatedAt, data.LastLogin, data.UpdatedAt = t, t, t
		query := composeInsert(i.TableName, data)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, "Error while executing : "+query)
		}
		bar.Add(1)
	}
}

func (i QueryParam) VisitorSeeder() {
	bar := progressbar.Default(200)
	data := Visitor{}
	t := time.Now().Format(dmyhms)
	rand.Seed(time.Now().UnixNano())
	for c := 0; c < 200; c++ {
		data.Id = c + 1
		if c < 120 {
			data.UserId = rand.Intn(44) + 6
			data.GuestId = 0
		} else {
			data.UserId = 0
			data.GuestId = c - 120 + 1
		}
		data.LoginAt = t
		data.Method = strconv.Itoa(rand.Intn(4))
		data.Purpose = gofakeit.Sentence(2)
		query := composeInsert(i.TableName, data)
		err := runQuery(i.Db, query)
		if err != nil {
			errFatal(err, "Error while executing : "+query)
		}
		bar.Add(1)
	}
}
