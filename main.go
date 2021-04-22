package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/dk"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	year := flag.Int("year", time.Now().Year(), "The year you want to calculate")
	workingHours := flag.Float64("hours", 7.5, "Configure the daily working hours")

	flag.Parse()

	for i := 1; i <= 12; i++ {
		hrs, err := calcWorkingHours(*year, i, *workingHours)
		if err != nil {
			return err
		}

		fmt.Printf("Total working hours in %02d/%d: %f \n", i, *year, hrs)
	}

	return nil
}

func calcWorkingHours(year, month int, workingHours float64) (float64, error) {
	c := cal.NewBusinessCalendar()
	c.AddHoliday(dk.Holidays...)

	firstDay, err := time.Parse("2006-1-02", fmt.Sprintf("%d-%d-01", year, month))
	lastDay := firstDay.AddDate(0, 1, 0).AddDate(0, 0, -1)
	if err != nil {
		return 0, err
	}

	days := c.WorkdaysInRange(firstDay, lastDay)

	hrs := float64(days) * workingHours

	return hrs, nil
}
