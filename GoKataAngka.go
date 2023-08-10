package main

import (
	"errors"
	"strings"
	"fmt"
	"strconv"
)

var numberSystem = map[string]int{
	"nol":     0,
	"kosong":  0,
	"pertama": 1,
	"satu":    1,
	"dua":     2,
	"tiga":    3,
	"empat":   4,
	"lima":    5,
	"enam":    6,
	"tujuh":   7,
	"delapan": 8,
	"sembilan": 9,
	"sepuluh":  10,
	"sebelas":  11,
	"belas":    10,
	"puluh":    10,
	"seratus":  100,
	"ratus":    100,
	"seribu":   1000,
	"sejuta":   1000000,
	"ribu":     1000,
	"juta":     1000000,
	"miliar":   1000000000,
	"milyar":   1000000000,
}

var numberSystemFloat = map[string]float64{
	"setengah": 0.5,
}
var multiplierSystem = []string{"puluh", "ratus", "ribu", "juta", "milyar", "trilyun", "miliar"}

func wordToNum(numberSentence string) (int, error) {
	numberSentence = strings.ReplaceAll(numberSentence, "-", " ")
	numberSentence = strings.ToLower(numberSentence)

	if strings.TrimSpace(numberSentence) == "" {
		return 0, errors.New("Input string is empty")
	}

	if _, err := strconv.Atoi(numberSentence); err == nil {
		return strconv.Atoi(numberSentence)
	}

	splitWords := strings.Fields(numberSentence)
	cleanNumbers := []string{}
	cleanDecimalNumbers := []int{}

	for _, word := range splitWords {
		if strings.HasPrefix(word, "ke") {
			word = strings.TrimPrefix(word, "ke")
		}
		if _, exists := numberSystem[word]; exists {
			cleanNumbers = append(cleanNumbers, word)
		}
		
		if _, exists := numberSystemFloat[word]; exists {
			cleanNumbers = append(cleanNumbers, word)
		}
	}

	if len(cleanNumbers) == 0 {
		return 0, errors.New("No valid number words found")
	}

	tmpNumber := 1
	biggestMultiplier := 1

	for i := len(cleanNumbers) - 1; i >= 0; i-- {
		w := cleanNumbers[i]
		if contains(multiplierSystem, w) {
			if numberSystem[w] > biggestMultiplier {
				biggestMultiplier = numberSystem[w]
				tmpNumber *= numberSystem[w]
			} else {
				tmpNumber = numberSystem[w] * biggestMultiplier
			}
		} else {
			if tmpNumber == 1 {
				cleanDecimalNumbers = append(cleanDecimalNumbers, numberSystem[w])
			} else {
				tmpNumber *= numberSystem[w]
				cleanDecimalNumbers = append(cleanDecimalNumbers, tmpNumber)
				tmpNumber = 1
			}
		}
	}

	negWords := []string{"-", "minus", "negatif"}
	for _, w := range negWords {
		if strings.Contains(numberSentence, w) {
			sum := 0
			for _, num := range cleanDecimalNumbers {
				sum += num
			}
			return sum * -1, nil
		}
	}

	sum := 0
	for _, num := range cleanDecimalNumbers {
		sum += num
	}
	return sum, nil
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func main() {
	result, err := wordToNum("seratus dua puluh lima setengah")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
