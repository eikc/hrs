package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	yearStr := os.Args[1]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return err
	}

	for i := 1; i <= 12; i++ {
		hrs, err := calcWorkingHours(year, i)
		if err != nil {
			return err
		}

		fmt.Printf("Total working hours in %d : %f \n", i, hrs)
	}

	return nil
}

func calcWorkingHours(year, month int) (float64, error) {
	c := cal.NewBusinessCalendar()
	c.AddHoliday(dk.Holidays...)

	firstDay, err := time.Parse("2006-1-02", fmt.Sprintf("%d-%d-01", year, month))
	lastDay := firstDay.AddDate(0, 1, 0).AddDate(0, 0, -1)
	if err != nil {
		return 0, err
	}

	days := c.WorkdaysInRange(firstDay, lastDay)

	hrs := float64(days) * 7.5

	return hrs, nil
}
