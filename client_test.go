package main

import "testing"
import "time"

func TestCurrentRFC1123FormattedDate(t *testing.T) {
	date := time.Date(2018, 04, 06, 22, 35, 0, 0, time.UTC)
	expected := "Fri, 06 Apr 2018 22:35:00 GMT"

	formatted := currentRFC1123FormattedDate(date)

	if formatted != expected {
		t.Errorf("Incorrectly formatted:\n got: %v\nwant: %v", formatted, expected)
	}
}
