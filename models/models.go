package models

import "time"

type User struct {
	ID   int
	TgID int64 `db:"tg_id"`
}

type TodoList struct {
	ID          int
	UserID      int `db:"user_id"`
	Title       string
	Description string
	CreatedAt   time.Time `db:"created_at"`
}

type Task struct {
	ID          int
	ListID      int `db:"list_id"`
	Title       string
	Description string
	DueDate     time.Time `db:"due_date"`
	IsDone      bool      `db:"is_done"`
	CreatedAt   time.Time `db:"created_at"`
}
