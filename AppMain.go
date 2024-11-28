package main

import (
	"database/sql"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	_ "github.com/lib/pq"
)

const (
	dbUser     = "your_db_user"     // Имя пользователя PostgreSQL
	dbPassword = "your_db_password" // Пароль PostgreSQL
	dbName     = "your_db_name"     // Имя базы данных PostgreSQL
)

func main() {
	// Инициализация приложения Fyne
	myApp := app.New()
	mainWindow := myApp.NewWindow("Приветствие")
	mainWindow.Resize(fyne.NewSize(400, 300))

	// Установка серого фона
	mainWindow.SetContent(container.NewVBox(
		widget.NewLabel("Добро пожаловать!"),
		widget.NewButton("Введите имя и пароль", func() {
			openCredentialsWindow(myApp)
		}),
	))

	mainWindow.ShowAndRun()
}

func openCredentialsWindow(myApp fyne.App) {
	// Окно для ввода имени и пароля
	credentialsWindow := myApp.NewWindow("Ввод имени и пароля")
	credentialsWindow.Resize(fyne.NewSize(300, 200))

	usernameEntry := widget.NewEntry()
	usernameEntry.SetPlaceHolder("Введите имя пользователя")

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.SetPlaceHolder("Введите пароль")

	submitButton := widget.NewButton("Сохранить", func() {
		username := usernameEntry.Text
		password := passwordEntry.Text

		if username == "" || password == "" {
			dialog.ShowInformation("Ошибка", "Имя пользователя и пароль не могут быть пустыми", credentialsWindow)
			return
		}

		err := saveToDatabase(username, password)
		if err != nil {
			dialog.ShowError(err, credentialsWindow)
		} else {
			dialog.ShowInformation("Успех", "Данные сохранены успешно!", credentialsWindow)
			credentialsWindow.Close()
		}
	})

	credentialsWindow.SetContent(container.NewVBox(
		widget.NewLabel("Введите ваши данные:"),
		usernameEntry,
		passwordEntry,
		submitButton,
	))

	credentialsWindow.Show()
}

func saveToDatabase(username, password string) error {
	// Подключение к базе данных PostgreSQL
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}
	defer db.Close()

	// Проверка подключения
	if err = db.Ping(); err != nil {
		return fmt.Errorf("ошибка проверки соединения с базой данных: %w", err)
	}

	// Вставка данных в базу
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err = db.Exec(query, username, password)
	if err != nil {
		return fmt.Errorf("ошибка сохранения данных: %w", err)
	}

	return nil
}
