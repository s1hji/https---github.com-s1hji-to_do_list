package gui

import (
	"time"
	"todolist/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const dateFormat = "02.01.2006"

func ShowTodoLists(w fyne.Window, lists []models.TodoList, items map[int][]models.Task) {
	listContainer := container.NewVBox()

	for _, todoList := range lists {
		currentList := todoList
		listButton := widget.NewButton(currentList.Title, func() {
			ShowTodoItems(w, currentList, items, lists)
		})
		listContainer.Add(listButton)
	}

	addListButton := widget.NewButton("+ Добавить список", func() {
		entry := widget.NewEntry()
		entry.SetPlaceHolder("Название списка")

		dialog.ShowForm(
			"Новый список задач",
			"Создать",
			"Отмена",
			[]*widget.FormItem{{Text: "", Widget: entry}},
			func(b bool) {
				if b && entry.Text != "" {
					newList := models.TodoList{
						ID:        len(lists) + 1,
						UserID:    1,
						Title:     entry.Text,
						CreatedAt: time.Now(),
					}
					lists = append(lists, newList)
					items[newList.ID] = []models.Task{}
					ShowTodoLists(w, lists, items)
				}
			},
			w,
		)
	})

	content := container.NewVBox(
		widget.NewLabelWithStyle("My Tasks", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		listContainer,
		addListButton,
	)

	w.SetContent(content)
}

func ShowTodoItems(w fyne.Window, todoList models.TodoList, items map[int][]models.Task, lists []models.TodoList) {
	listItems := items[todoList.ID]
	tasksContainer := container.NewVBox()

	for i := range listItems {
		item := &listItems[i]

		taskBtn := widget.NewButton("", nil)
		taskBtn.Alignment = widget.ButtonAlignLeading

		updateTaskButton := func() {
			taskText := item.Title
			if !item.DueDate.IsZero() {
				taskText += " (" + item.DueDate.Format("02.01") + ")"
			}
			taskBtn.SetText(taskText)
		}
		updateTaskButton()

		taskBtn.OnTapped = func() {
			showTaskDetails(w, item, updateTaskButton)
		}

		check := widget.NewCheck("", func(done bool) {
			item.IsDone = done
			updateTaskButton()
		})
		check.SetChecked(item.IsDone)

		deleteBtn := widget.NewButton("✕", func() {
			dialog.ShowConfirm(
				"Удаление задачи",
				"Удалить эту задачу?",
				func(b bool) {
					if b {
						newItems := append(listItems[:i], listItems[i+1:]...)
						items[todoList.ID] = newItems
						ShowTodoItems(w, todoList, items, lists)
					}
				},
				w,
			)
		})
		deleteBtn.Importance = widget.LowImportance

		taskRow := container.NewHBox(
			check,
			taskBtn,
			layout.NewSpacer(),
			deleteBtn,
		)

		tasksContainer.Add(taskRow)
	}

	addButton := widget.NewButton("+ Добавить задачу", func() {
		titleEntry := widget.NewEntry()
		titleEntry.SetPlaceHolder("Название задачи")

		descEntry := widget.NewEntry()
		descEntry.SetPlaceHolder("Описание")
		descEntry.MultiLine = true

		dateEntry := widget.NewEntry()
		dateEntry.SetPlaceHolder("дд.мм.гггг")

		dialog.ShowForm(
			"Новая задача",
			"Добавить",
			"Отмена",
			[]*widget.FormItem{
				{Text: "Название:", Widget: titleEntry},
				{Text: "Описание:", Widget: descEntry},
				{Text: "Срок выполнения:", Widget: dateEntry},
			},
			func(b bool) {
				if !b {
					return
				}

				var dueDate time.Time
				if dateEntry.Text != "" {
					parsedDate, err := time.Parse("02.01.2006", dateEntry.Text)
					if err == nil {
						dueDate = parsedDate
					}
				}

				newTask := models.Task{
					ID:          len(items[todoList.ID]) + 1,
					ListID:      todoList.ID,
					Title:       titleEntry.Text,
					Description: descEntry.Text,
					DueDate:     dueDate,
					IsDone:      false,
					CreatedAt:   time.Now(),
				}

				items[todoList.ID] = append(items[todoList.ID], newTask)
				ShowTodoItems(w, todoList, items, lists)
			},
			w,
		)
	})

	backButton := widget.NewButton("← Назад", func() {
		ShowTodoLists(w, lists, items)
	})

	content := container.NewVBox(
		widget.NewLabelWithStyle(todoList.Title, fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabel(todoList.Description),
		tasksContainer,
		addButton,
		backButton,
	)

	w.SetContent(content)
}


func showTaskDetails(w fyne.Window, task *models.Task, onUpdate func()) {
	titleLabel := widget.NewLabelWithStyle(task.Title, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	descLabel := widget.NewLabel(task.Description)
	descLabel.Wrapping = fyne.TextWrapWord

	var dateText string
	if !task.DueDate.IsZero() {
		dateText = "Срок выполнения: " + task.DueDate.Format("02.01.2006")
	} else {
		dateText = "Срок не установлен"
	}
	dateLabel := widget.NewLabel(dateText)

	editBtn := widget.NewButton("Редактировать", func() {
		editTaskDialog(w, task, func() {
			onUpdate()
			showTaskDetails(w, task, onUpdate)
		})
	})

	content := container.NewVBox(
		titleLabel,
		widget.NewSeparator(),
		descLabel,
		dateLabel,
		layout.NewSpacer(),
		editBtn,
	)

	dialog.ShowCustom("Детали задачи", "Закрыть", content, w)
}


func editTaskDialog(w fyne.Window, task *models.Task, onSave func()) {
	titleEntry := widget.NewEntry()
	titleEntry.SetText(task.Title)

	descEntry := widget.NewEntry()
	descEntry.SetText(task.Description)
	descEntry.MultiLine = true

	dateEntry := widget.NewEntry()
	if !task.DueDate.IsZero() {
		dateEntry.SetText(task.DueDate.Format(dateFormat))
	}
	dateEntry.SetPlaceHolder(dateFormat)

	dialog.ShowForm(
		"Редактировать задачу",
		"Сохранить",
		"Отмена",
		[]*widget.FormItem{
			{Text: "Название:", Widget: titleEntry},
			{Text: "Описание:", Widget: descEntry},
			{Text: "Срок выполнения:", Widget: dateEntry},
		},
		func(b bool) {
			if !b {
				return
			}

			task.Title = titleEntry.Text
			task.Description = descEntry.Text

			if dateEntry.Text != "" {
				parsedDate, err := time.Parse(dateFormat, dateEntry.Text)
				if err == nil {
					task.DueDate = parsedDate
				}
			} else {
				task.DueDate = time.Time{}
			}

			if onSave != nil {
				onSave()
			}
		},
		w,
	)
}
