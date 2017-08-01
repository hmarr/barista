package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TimeRange struct {
	Start time.Duration
	End   time.Duration
}

func (t *TimeRange) UnmarshalJSON(data []byte) error {
	var strRange string
	if err := json.Unmarshal(data, &strRange); err != nil {
		return err
	}

	parsedRange, err := parseTimeRange(strRange)
	if err != nil {
		return err
	}
	t.Start = parsedRange.Start
	t.End = parsedRange.End
	return nil
}

func parseTimeRange(r string) (*TimeRange, error) {
	parts := strings.Split(r, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Invalid range %s", r)
	}

	startTime, err := time.Parse("15:04", strings.TrimSpace(parts[0]))
	if err != nil {
		return nil, err
	}

	endTime, err := time.Parse("15:04", strings.TrimSpace(parts[1]))
	if err != nil {
		return nil, err
	}

	start := time.Duration(startTime.Hour()) * time.Hour
	start += time.Duration(startTime.Minute()) * time.Minute

	end := time.Duration(endTime.Hour()) * time.Hour
	end += time.Duration(endTime.Minute()) * time.Minute

	return &TimeRange{Start: start, End: end}, nil
}
