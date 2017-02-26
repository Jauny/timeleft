package timeleft

import (
	"fmt"
	"time"
)

const (
	dateLayout  = "2006-01-02"
	hoursPerDay = 24
)

func Countdown(d string) (time.Duration, error) {
	var left time.Duration

	then, err := time.Parse(dateLayout, d)
	if err != nil {
		return left, err
	}

	left = then.Sub(time.Now())
	return left, nil
}

func FormatDuration(d time.Duration) string {
	totalHours := int64(d / time.Hour)
	daysLeft := totalHours / hoursPerDay
	hoursLeft := totalHours % hoursPerDay

	return fmt.Sprintf("You have %v days and %v hours left!", daysLeft, hoursLeft)
}
