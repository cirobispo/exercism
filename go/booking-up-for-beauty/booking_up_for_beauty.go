package booking

import (
	"fmt"
	"strings"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	dp, e := time.Parse("1/02/2006 15:04:05", date)
	if e == nil {
		return dp
	}
	panic(fmt.Sprintf("date parse error: %v", e.Error()))
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	dp, _ := time.Parse("January 2, 2006 15:04:05", date)
	yl := dp.Year() - time.Now().Year()
	ml := dp.Month() - time.Now().Month()
	dl := dp.Day() - time.Now().Day()

	if yl == 0 { // 2022 == 2022
		if ml == 0 {
			return dl < 0
		}
		return ml < 0
	}
	return yl < 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	//Friday, March 8, 1974 12:02:02
	dp, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	return dp.Hour() >= 12 && dp.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	dt := parseDate(date)
	tm := parseTime(date)

	result := fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.", dt.Weekday(), dt.Month(), dt.Day(), dt.Year(), tm.Hour(), tm.Minute())
	return result
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}

func parseDate(date string) time.Time {
	if dateparts := strings.Split(date, " "); len(dateparts) > 0 {
		var format string
		dmy := strings.Split(dateparts[0], "/")
		if len(dmy) == 3 {
			if len(dmy[0]) == 2 {
				format += "0"
			}
			format += "1/"
			if len(dmy[1]) == 2 {
				format += "0"
			}
			format += "2/"
			if len(dmy[2]) == 4 {
				format += "20"
			}
			format += "06"
		}

		dp, _ := time.Parse(format, dateparts[0])
		return dp
	}
	panic(" parse error")
}

func parseTime(date string) time.Time {
	if dateparts := strings.Split(date, " "); len(dateparts) > 1 {
		dp, _ := time.Parse("15:04:05", dateparts[1])
		return dp
	}
	panic(" parse error")
}
