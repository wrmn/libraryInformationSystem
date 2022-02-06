package main

import (
	"fmt"
	"strconv"

	"github.com/bxcodec/faker/v3"
)

//make query for insert data user
func (i InsertParam) DdcSeeder() {
	ddc := Ddcs{}
	var num string
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
	}

}

//make query for insert data user
func (i InsertParam) UserSeeder() {
	fmt.Println("called this")
}
