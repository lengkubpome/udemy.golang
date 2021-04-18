package report

import (
	"fmt"
	"strings"
	"testing"
)

func TestBasicReport(t *testing.T) {

	getReport = func(id string) string {
		// connect to database
		return "xxxxx"
	}

	report := generateReport("12DB")
	if !strings.Contains(report, ": VERIFIED.") {
		t.Error("failed")
	}
	fmt.Println("got report : ", report)
}
