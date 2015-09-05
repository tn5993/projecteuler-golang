package calendar

import (
	"github.com/tn5993/projecteuler/utils/simplemath/intmath"
)

func IsLeapYear(year int) bool {
	return year%100 != 0 && year%4 == 0 || year%400 == 0
}

func GetDayOfWeek(month, day, fullyear int) int {
	doomday := GetDoomDayOfYear(fullyear)
	doomMap := getDoomMap(fullyear)
	diff := intmath.Abs(doomMap[month] - day)
	if doomMap[month] >= day {
		res := (doomday - diff) % 7
		if res < 0 {
			res += 7
		}
		return res
	}

	return (diff + doomday) % 7
}

/*
	Sunday = 0
	Monday = 1
	Tuesday = 2
	Wednesday = 3
	Thursday = 4
	Friday = 5
	Saturday = 6
*/
func GetDoomDayOfYear(fullyear int) int {
	year := fullyear % 100 //get last 2 digit of year
	anchor_day := getAnchorDay(fullyear)
	return (year/12 + year%12 + (year%12)/4 + anchor_day) % 7
}

func getDoomMap(fullyear int) map[int]int {
	m := map[int]int{3: 0, 4: 4, 5: 9, 6: 6, 7: 11, 8: 8, 9: 5, 10: 10, 11: 7, 12: 12}
	if IsLeapYear(fullyear) {
		m[1] = 4
		m[2] = 29
	} else {
		m[1] = 3
		m[2] = 28
	}
	return m
}

func getAnchorDay(fullyear int) int {
	table := []int{2, 0, 5, 3}
	return table[(fullyear/100)%4]
}
