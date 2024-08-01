package romannumerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

var cases = []struct {
	Arabic uint16
	Roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{90, "XC"},
	{100, "C"},
	{1984, "MCMLXXXIV"},
}

func TestRomanNumerals(t *testing.T) {
	// t.Run("1 gets converted to I", func(t *testing.T) {

	// 	got := ConvertToRoman(1)
	// 	want := "I"

	// 	if got != want {
	// 		t.Errorf("got %q, want %q", got, want)
	// 	}
	// })
	// t.Run("2 gets converted to II", func(t *testing.T) {
	// 	got := ConvertToRoman(2)
	// 	want := "II"
	// 	if got != want {
	// 		t.Errorf("got %q, want %q", got, want)
	// 	}
	// })

	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}

}

func TestConvertingToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}
}
