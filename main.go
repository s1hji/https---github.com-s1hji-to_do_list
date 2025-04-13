package main

import (
	"time"
	"todolist/gui"
	"todolist/models"
	"todolist/theme"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&theme.CustomTheme{})

	w := a.NewWindow("My Tasks")
	w.Resize(fyne.NewSize(400, 600))

	// Тестовые данные с учетом новых полей
	lists := []models.TodoList{
		{
			ID:          1,
			UserID:      1,
			Title:       "Задачи на 2025 год",
			Description: "Годовые цели и задачи",
			CreatedAt:   time.Now(),
		},
		{
			ID:          2,
			UserID:      1,
			Title:       "Задачи на март",
			Description: "Планы на март",
			CreatedAt:   time.Now(),
		},
	}

	tasks := map[int][]models.Task{
		1: {
			{
				ID:          1,
				ListID:      1,
				Title:       "Go shopping",
				Description: "Купить продукты",
				DueDate:     time.Now().AddDate(0, 0, 3),
				IsDone:      false,
				CreatedAt:   time.Now(),
			},
			{
				ID:          2,
				ListID:      1,
				Title:       "Short exercise",
				Description: "15 минут упражнений",
				DueDate:     time.Now().AddDate(0, 0, 1),
				IsDone:      false,
				CreatedAt:   time.Now(),
			},
		},
	}

	gui.ShowTodoLists(w, lists, tasks)
	w.ShowAndRun()
}
