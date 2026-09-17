package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

const schema = `
CREATE TABLE scheduler (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	date CHAR(8) NOT NULL DEFAULT '',
	title VARCHAR NOT NULL DEFAULT '',
	comment TEXT NOT NULL DEFAULT '',
	repeat VARCHAR(128) NOT NULL DEFAULT ''
); 
`

const createIndex = `
CREATE INDEX idx_date ON scheduler(date);
`

var db *sql.DB
var ErrNotFound = errors.New("not found")

// Init открывает базу данных и создаёт необходимые таблицы и индексы,
// если файл базы данных отсутствует.
func Init(dbFile string) error {
	_, err := os.Stat(dbFile)

	var install bool

	if err != nil {
		if os.IsNotExist(err) {
			install = true
		} else {
			return fmt.Errorf("check database file: %w", err)
		}
	}

	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	//проверить
	if err := db.Ping(); err != nil {
		db.Close()
		return fmt.Errorf("ping database: %w", err)
	}

	if install {
		_, err = db.Exec(schema)
		if err != nil {
			return fmt.Errorf("create table: %w", err)
		}
		_, err = db.Exec(createIndex)
		if err != nil {
			return fmt.Errorf("create index: %w", err)
		}
	}
	return nil
}

// GetTask возвращает задачу по её идентификатору.
func GetTask(id string) (*Task, error) {
	query := `
	SELECT id, date, title, comment, repeat
	FROM scheduler
	WHERE id = ?`

	task := &Task{}

	err := db.QueryRow(query, id).Scan(
		&task.ID,
		&task.Date,
		&task.Title,
		&task.Comment,
		&task.Repeat,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("get task: %w", err)
	}
	return task, nil
}

// Close закрывает соединение с базой данных.
func Close() error {
	if db == nil {
		return nil
	}
	return db.Close()
}
