package api

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const dateLayout = "20060102"

func afterNow(date, now time.Time) bool {
	return date.After(now)
}

// NextDate вычисляет следующую дату задачи по правилу повторения
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("empty repeat")
	}

	date, err := time.Parse(dateLayout, dstart)
	if err != nil {
		return "", fmt.Errorf("parse start date: %w", err)
	}

	parts := strings.Fields(repeat)
	if len(parts) == 0 {
		return "", errors.New("invalid repeat rule")
	}

	switch parts[0] {

	case "d":
		if len(parts) != 2 {
			return "", errors.New("invalid day repeat format")
		}
		interval, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf("parse day interval: %w", err)
		}

		if interval < 1 || interval > 400 {
			return "", fmt.Errorf("invalid day interval: %d", interval)
		}

		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format(dateLayout), nil
	case "y":
		if len(parts) != 1 {
			return "", errors.New("invalid yearly repeat format")
		}

		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}
		return date.Format(dateLayout), nil

	default:
		return "", errors.New("unsupported repeat rule")
	}
}
