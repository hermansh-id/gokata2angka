package main

import (
	"testing"
)

func TestWordToNum(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"nol", 0},
		{"satu", 1},
		{"dua", 2},
		{"sembilan", 9},
		{"sepuluh", 10},
		{"sebelas", 11},
		{"dua puluh", 20},
		{"dua puluh satu", 21},
		{"lima puluh", 50},
		{"seratus", 100},
		{"seratus dua puluh lima", 125},
		{"seribu", 1000},
		{"seribu lima ratus", 1500},
		{"sejuta", 1000000},
		{"dua juta tiga ratus ribu empat ratus lima puluh enam", 2300456},
		{"sembilan juta delapan ratus tiga puluh dua ribu seratus lima puluh enam", 9832156},
		// ... Add more test cases ...
	}

	for _, test := range tests {
		result, err := wordToNum(test.input)
		if err != nil {
			t.Errorf("For input '%s', error was not expected but got: %v", test.input, err)
		}
		if result != test.expect {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expect, result)
		}
	}
}

func TestAdditionalCases(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"kosong", 0},
		{"minus dua", -2},
		{"negatif sepuluh", -10},
		{"pertama", 1},
		{"kedua belas", 12},
		{"dua puluh ketiga", 23},
		{"seratus dua puluh kelima", 125},
		{"seribu lima ratus keenam", 1506},
		// ... Add more test cases ...
	}

	for _, test := range tests {
		result, err := wordToNum(test.input)
		if err != nil {
			t.Errorf("For input '%s', error was not expected but got: %v", test.input, err)
		}
		if result != test.expect {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expect, result)
		}
	}
}
