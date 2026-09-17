package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Detaleryn/final/pkg/db"
)

func addTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task db.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	if task.Title == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "title is required",
		})
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		log.Printf("failed to add task: %v", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"id": strconv.FormatInt(id, 10),
	})

}

// checkDate проверяет дату и правило повторения задачи,
// а для просроченной задачи вычисляет новую дату
func checkDate(task *db.Task) error {
	now := time.Now()

	if task.Date == "" {
		task.Date = now.Format(dateLayout)
	}

	t, err := time.Parse(dateLayout, task.Date)
	if err != nil {
		return fmt.Errorf("invalid date format, expected YYYYMMDD: %w", err)
	}

	if task.Repeat != "" {
		if _, err := NextDate(now, task.Date, task.Repeat); err != nil {
			return fmt.Errorf("invalid repeat rule: %w", err)
		}
	}

	today, _ := time.Parse(dateLayout, now.Format(dateLayout))

	if t.Before(today) {
		if task.Repeat == "" {
			task.Date = today.Format(dateLayout)
		} else {
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return fmt.Errorf("invalid repeat rule: %w", err)
			}

			task.Date = next
		}
	}
	return nil
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode JSON response: %v", err)
	}
}
