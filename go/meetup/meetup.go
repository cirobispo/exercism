package meetup

import (
	"time"
)

// Define the WeekSchedule type here.
type WeekSchedule int

const (
	First WeekSchedule = 1
	Second WeekSchedule = 8
	Third WeekSchedule = 15
	Fourth WeekSchedule = 22
	Teenth WeekSchedule = 13
	Last WeekSchedule = -1
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	var result time.Time
	if wSched == Last {
		result=time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)
		result=result.AddDate(0, 0, -1)
	} else {
		result=time.Date(year, month, int(wSched), 0, 0, 0, 0, time.UTC)
	}

	if wSched != Last {
		currentWD:=result.Weekday()
		if currentWD > wDay {
			result=result.AddDate(0, 0, 7 - (int(currentWD) - int(wDay)))
		} else {
			result=result.AddDate(0, 0, int(wDay) - int(currentWD))
		}
	} else {
		currentWD:=result.Weekday()
		if currentWD >= wDay {
			result=result.AddDate(0, 0, int(wDay) - int(currentWD))
		} else {
			result=result.AddDate(0, 0, -7 + (int(wDay) - int(currentWD)))
		}
	}

	return result.Day()
}
