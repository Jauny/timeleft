package timeleft

import (
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	_, err := Countdown("2018-02-17")
	if err != nil {
		t.Errorf("Failed with unknown error.")
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
