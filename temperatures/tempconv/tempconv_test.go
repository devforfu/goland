package tempconv

import (
	"testing"
)

func TestCelsiusFahrenheitConversion(t *testing.T) {
	var testCases = []struct {
		inputValue float64
		expected float64
	}{
		{0,32},
		{100, 212},
		{-100, -148},
	}
	for _, testCase := range testCases {
		celsius := Celsius(testCase.inputValue)
		fahrenheit := Fahrenheit(testCase.expected)
		backAndForth := FToC(CToF(celsius))
		if backAndForth != celsius {
			t.Error("FToC(CToF(celsius)) != celsius")
		}
		if CToF(celsius) != fahrenheit {
			t.Error("CToF(celsius) != fahrenheit")
		}
		if FToC(fahrenheit) != celsius {
			t.Error("FToC(fahrenheit) != celsius")
		}
	}
}
