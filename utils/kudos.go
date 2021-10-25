package utils

var POINTS_TO_KUDOS = map[int]string{
	5:   "OK",
	10:  "NICE",
	20:  "GOOD",
	50:  "GREAT",
	100: "SUPER",
}

var KUDOS_TO_REAL = map[string]int{
	"OK":    2,
	"NICE":  5,
	"GOOD":  8,
	"GREAT": 15,
	"SUPER": 25,
}
