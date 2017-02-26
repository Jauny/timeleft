package timeleft

import (
	"time"
)

const dateLayout = "2006-01-02"

func Countdown(d string) (time.Duration, error) {
	var left time.Duration

	then, err := time.Parse(dateLayout, d)
	if err != nil {
		return left, err
	}

	left = then.Sub(time.Now())
	return left, nil
}
