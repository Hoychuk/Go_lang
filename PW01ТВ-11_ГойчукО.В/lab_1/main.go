package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"go1/handlers"
)

// Обробник для головної сторінки
func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Помилка завантаження сторінки", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Налаштування маршрутів
	http.HandleFunc("/", homePage)
	http.HandleFunc("/calculator", handlers.TaskHandler)

	// Запуск сервера
	port := ":8080"
	fmt.Println("Сервер запущено на http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
