package utils

var NumberToWord = map[int]string{
	1:  "um",
	2:  "dois",
	3:  "três",
	4:  "quatro",
	5:  "cinco",
	6:  "seis",
	7:  "sete",
	8:  "oito",
	9:  "nove",
	10: "dez",
	11: "onze",
	12: "doze",
	13: "treze",
	14: "quatroze",
	15: "quinze",
	16: "dezesseis",
	17: "dezessete",
	18: "dezoito",
	19: "dezenove",
	20: "vinte",
	30: "trinta",
	40: "quarenta",
	50: "cinquenta",
	60: "sessenta",
	70: "setenta",
	80: "oitenta",
	90: "noventa",
}

func convert1to99(n int) (w string) {
	if n < 20 {
		w = NumberToWord[n]
		return
	}

	r := n % 10
	if r == 0 {
		w = NumberToWord[n]
	} else {
		w = NumberToWord[n-r] + " e " + NumberToWord[r]
	}
	return
}

func convert100to999(n int) (w string) {
	q := n / 100
	r := n % 100

	if NumberToWord[q] == "um" {
		w = "cento "
	} else if NumberToWord[q] == "dois" {
		w = "duzentos "
	} else if NumberToWord[q] == "três" {
		w = "trezentos "
	} else if NumberToWord[q] == "quatro" {
		w = "quatrocentos "
	} else if NumberToWord[q] == "cinco" {
		w = "quinhentos "
	} else if NumberToWord[q] == "seis" {
		w = "seiscentos "
	} else if NumberToWord[q] == "sete" {
		w = "setessentos "
	} else if NumberToWord[q] == "oito" {
		w = "oitocentos "
	} else if NumberToWord[q] == "nove" {
		w = "novecentos "
	}

	if r == 0 {
		if w == "cento " {
			w = "cem"
		}
		return
	} else {
		w = w + "e " + convert1to99(r)
	}
	return
}

func convert1000to99999(n int) (w string) {
	q := n / 1000
	r := n % 1000

	if NumberToWord[q] == "um" {
		if r%100 == 0 && r != 0 {
			w = "mil e"
		} else {
			w = "mil"
		}
	} else {
		if r%100 == 0 && r != 0 {
			if q >= 100 {
				c := q / 100

				if NumberToWord[c] == "um" {
					w = "cem "
				} else if NumberToWord[c] == "dois" {
					w = "duzentos "
				} else if NumberToWord[c] == "três" {
					w = "trezentos "
				} else if NumberToWord[c] == "quatro" {
					w = "quatrocentos "
				} else if NumberToWord[c] == "cinco" {
					w = "quinhentos "
				} else if NumberToWord[c] == "seis" {
					w = "seiscentos "
				} else if NumberToWord[c] == "sete" {
					w = "setessentos "
				} else if NumberToWord[c] == "oito" {
					w = "oitocentos "
				} else if NumberToWord[c] == "nove" {
					w = "novecentos "
				}

				w += "mil e"
			} else {
				d := q / 10
				u := q % 10
				w = NumberToWord[d*10] + " e " + NumberToWord[u] + " mil e"
			}
		} else {
			if q > 20 && q%10 != 0 {
				d := q / 10
				u := q % 10

				w = NumberToWord[d*10] + " e " + NumberToWord[u] + " mil"

				if q > 100 {
					c := q / 100

					if NumberToWord[c] == "um" {
						w = "cento "
					} else if NumberToWord[c] == "dois" {
						w = "duzentos "
					} else if NumberToWord[c] == "três" {
						w = "trezentos "
					} else if NumberToWord[c] == "quatro" {
						w = "quatrocentos "
					} else if NumberToWord[c] == "cinco" {
						w = "quinhentos "
					} else if NumberToWord[c] == "seis" {
						w = "seiscentos "
					} else if NumberToWord[c] == "sete" {
						w = "setessentos "
					} else if NumberToWord[c] == "oito" {
						w = "oitocentos "
					} else if NumberToWord[c] == "nove" {
						w = "novecentos "
					}

					w += "e " + NumberToWord[u] + " mil"
				}

			} else {
				if q >= 100 {
					c := q / 100

					if NumberToWord[c] == "um" {
						w = "cem "
					} else if NumberToWord[c] == "dois" {
						w = "duzentos "
					} else if NumberToWord[c] == "três" {
						w = "trezentos "
					} else if NumberToWord[c] == "quatro" {
						w = "quatrocentos "
					} else if NumberToWord[c] == "cinco" {
						w = "quinhentos "
					} else if NumberToWord[c] == "seis" {
						w = "seiscentos "
					} else if NumberToWord[c] == "sete" {
						w = "setessentos "
					} else if NumberToWord[c] == "oito" {
						w = "oitocentos "
					} else if NumberToWord[c] == "nove" {
						w = "novecentos "
					}

					w += "mil"
				} else {
					w = NumberToWord[q] + " mil"
				}
			}
		}
	}

	if r == 0 {
		return
	} else {
		w = w + " " + convert100to999(r)
	}
	return
}

func Convert1to999999(n int) (w string) {
	if n > 999999 || n < 1 {
		panic("func Convert1to999999: n > 999999 or n < 1")
	}

	if n < 100 {
		w = convert1to99(n)
		return
	}
	if n < 1000 {
		w = convert100to999(n)
		return
	} else {
		w = convert1000to99999(n)
		return
	}
}

// func main() {
// 	for i := 1; i <= 999999; i++ {
// 		fmt.Println(Convert1to999999(i))
// 	}
// }
