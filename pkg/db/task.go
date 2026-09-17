package db

import (
	"fmt"
)

// Task представляет задачу планировщика.
type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// AddTask добавляет новую задачу в БД
func AddTask(task *Task) (int64, error) {

	query := `
	INSERT INTO scheduler (date, title, comment, repeat)
	VALUES (?, ?, ?, ?) `

	res, err := db.Exec(
		query,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
	)
	if err != nil {
		return 0, fmt.Errorf("add task: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get inserted task ID: %w", err)
	}
	return id, nil
}

// Tasks возвращает список задач, отсортированных по дате
func Tasks(limit int) ([]*Task, error) {

	query := `
	SELECT id, date, title, comment, repeat
	FROM scheduler
	ORDER BY date
	LIMIT ?`

	args := []any{limit}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("get tasks: %w", err)
	}

	defer rows.Close()

	tasks := make([]*Task, 0)

	for rows.Next() {
		task := &Task{}

		err := rows.Scan(
			&task.ID,
			&task.Date,
			&task.Title,
			&task.Comment,
			&task.Repeat,
		)
		if err != nil {
			return nil, fmt.Errorf("scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate tasks: %w", err)
	}
	return tasks, nil
}

// UpdateTask обновляет данные существующей задачи
func UpdateTask(task *Task) error {
	query := `
	UPDATE scheduler
	SET 
		date = ?,
		title = ?,
		comment = ?,
		repeat = ?
	WHERE id = ?`

	res, err := db.Exec(
		query,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
		task.ID,
	)
	if err != nil {
		return fmt.Errorf("update task: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("incorrect id for updating task")
	}
	return nil
}

// UpdateDate изменяет дату существующей задачи
func UpdateDate(next string, id string) error {
	query := `
	UPDATE scheduler
	SET date = ?
	WHERE id = ?`

	res, err := db.Exec(
		query,
		next,
		id,
	)
	if err != nil {
		return fmt.Errorf("update date: %w", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("incorrect id for updating task date")
	}
	return nil
}

// DeleteTask удаляет задачу из базы данных
func DeleteTask(id string) error {
	query := `
	DELETE FROM scheduler
	WHERE id = ?`

	res, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("incorrect id for deleting task")
	}
	return nil
}
