package secondstostring

import (
	"math"
	"strconv"
	"strings"
)

func secondsToString(seconds int64) string {

	if seconds == 0 {
		return "now"
	}

	// holds the output as distinct elements to be joined at the end
	timeStrings := []string{}

	// maps haven't got a guarenteed iteration order, wtf. So I made a slice of structs
	timeUnits := []struct {
		value int64
		name  string
	}{
		{31536000, "year"},
		{86400, "day"},
		{3600, "hour"},
		{60, "minute"},
		{1, "second"},
	}

	for _, unit := range timeUnits {
		if seconds >= unit.value {
			timeStrings = append(timeStrings, generateOutput(seconds, unit.value, unit.name))
			seconds = seconds % unit.value
		}
	}

	if len(timeStrings) == 1 {
		return timeStrings[0]
	}

	if len(timeStrings) > 2 {
		return strings.Join(timeStrings[:len(timeStrings)-1], ", ") + " and " + timeStrings[len(timeStrings)-1]
	}

	return strings.Join(timeStrings, " and ")
}

func generateOutput(seconds int64, unit int64, unitName string) string {
	output := int64(math.Floor(float64(seconds) / float64(unit)))

	appender := ""
	if output > 1 {
		appender = "s"
	}

	return strconv.FormatInt(output, 10) + " " + unitName + appender
}
