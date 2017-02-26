package main

import (
	"flag"
	"fmt"
	"github.com/jauny/timeleft/timeleft"
	"os"
)

var dateUntil string

func initDate() {
	const (
		defaultDate = "2018-02-17"
		usage       = "The date to countdown to"
	)
	flag.StringVar(&dateUntil, "date", defaultDate, usage)
	flag.StringVar(&dateUntil, "d", defaultDate, usage+" (shorhand)")
	flag.Parse()
}

func main() {
	if err := do(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func do() error {
	initDate()
	left, err := timeleft.Countdown(dateUntil)
	if err != nil {
		return err
	}
	fmt.Println(left)

	return nil
}
