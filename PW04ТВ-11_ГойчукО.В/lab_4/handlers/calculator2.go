package handlers

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

// Функція для обробки запиту на калькулятор 2
func Task2Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		SNom := 6.3

		// Отримуємо значення з форми
		P, _ := strconv.ParseFloat(r.FormValue("P"), 64)
		Usn, _ := strconv.ParseFloat(r.FormValue("Usn"), 64)

		//Проводимо розрахунки згідно формул
		Xc := (Usn * Usn) / P
		Xt := (Usn * Usn * Usn) / SNom / 100

		X := Xc + Xt

		Ip0 := Usn / math.Sqrt(3.0) / X

		// Форматуємо числа до N знаків після коми
		precision := 2
		tmpl, _ := template.ParseFiles("templates/calculator2.html")
		tmpl.Execute(w, map[string]string{
			"Xc":  roundFloat(Xc, precision),
			"Xt":  roundFloat(Xt, precision),
			"X":   roundFloat(X, precision),
			"Ip0": roundFloat(Ip0, precision),
		})
		return
	}

	// Відображення HTML-сторінки
	tmpl, _ := template.ParseFiles("templates/calculator2.html")
	tmpl.Execute(w, nil)
}
