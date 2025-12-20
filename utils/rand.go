package utils

import (
	"math/rand"
	"strconv"
)

func Generate6DigitCode() string {
	code := rand.Intn(900000) + 100000
	return strconv.Itoa(code)
}
