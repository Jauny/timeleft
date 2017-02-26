package timeleft

import (
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	date := time.Now().Add(time.Hour * 50)
	dateString := date.String()[:10]
	dateParsed, _ := time.Parse("2006-01-02", dateString)
	expectedResult := dateParsed.Sub(time.Now())

	result, err := Countdown(dateString)
	if err != nil {
		t.Errorf("Failed %v.", err)
		return
	}

	if (result / time.Second) != (expectedResult / time.Second) {
		t.Errorf("Failed with result %v. Expected %v", result, expectedResult)
	}
}

func TestFormatDuration(t *testing.T) {
	left := time.Hour * 50
	expected := "You have 2 days and 2 hours left!"
	formatted := FormatDuration(left)

	if formatted != expected {
		t.Errorf("Failed with string %v. Expected %v", formatted, expected)
	}
}
