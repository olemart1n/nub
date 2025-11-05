package utils

import "time"

func HumanDate(t time.Time) string {
	return t.Format("02.01.2006")
}
