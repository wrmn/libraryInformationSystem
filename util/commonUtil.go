package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

// Format int to ddc format
func IntToDdc(c int) string {
	var num string
	if c < 10 {
		num = fmt.Sprintf("00%d", c)
	} else if c < 100 {
		num = fmt.Sprintf("0%d", c)
	} else {
		num = strconv.Itoa(c)
	}
	return num
}

// give string of random digit with n lenght
func RandDigit(n int) (s string) {
	var d int
	for c := 0; c < n; c++ {
		if c == 0 {
			d = rand.Intn(8) + 1
		} else {
			d = rand.Intn(9)
		}
		s = fmt.Sprintf("%s%d", s, d)
	}
	return s
}

// date random in time.Time with param string formatted date. format param : yyyy-mm-dd
func DateRandom(minYear string, maxYear string) time.Time {
	min, _ := time.Parse(Dmy, minYear)
	max, _ := time.Parse(Dmy, maxYear)
	return gofakeit.DateRange(min, max)
}
