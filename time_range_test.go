package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	timeRange, err := parseTimeRange("09:00 - 18:30")
	if err != nil {
		t.Error("Error parsing valid range")
		return
	}

	if timeRange.Start.String() != "9h0m0s" {
		t.Errorf("Incorrect start of range %s", timeRange.Start)
	}

	if timeRange.End.String() != "18h30m0s" {
		t.Errorf("Error parsing ned of range %s", timeRange.End)
	}
}

func TestParseMissingDash(t *testing.T) {
	_, err := parseTimeRange("09:00 18:30")
	if err == nil {
		t.Error("Expected error parsing invalid range")
	}
}

func TestParseTooManyParts(t *testing.T) {
	_, err := parseTimeRange("09:00 - 18:30 - 20:00")
	if err == nil {
		t.Error("Expected error parsing invalid range")
	}
}

func TestParseInvalidTime(t *testing.T) {
	_, err := parseTimeRange("09:00 - 25:30")
	if err == nil {
		t.Error("Expected error parsing invalid range")
	}
}
