package handlers

import (
	"html/template"
	"net/http"
	"strconv"
)

// Функція для обробки запиту на калькулятор 1
func TaskHandler2(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Отримуємо значення з форми
		zPerA, _ := strconv.ParseFloat(r.FormValue("zPerA"), 64)
		zPerP, _ := strconv.ParseFloat(r.FormValue("zPerP"), 64)
		omega, _ := strconv.ParseFloat(r.FormValue("omega"), 64)
		tV, _ := strconv.ParseFloat(r.FormValue("tV"), 64)
		Pm, _ := strconv.ParseFloat(r.FormValue("Pm"), 64)
		Tm, _ := strconv.ParseFloat(r.FormValue("Tm"), 64)
		kP, _ := strconv.ParseFloat(r.FormValue("kP"), 64)

		mWnedA := omega * tV * Pm * Tm
		mWnedP := kP * Pm * Tm
		mZper := zPerA*mWnedA + zPerP*mWnedP

		// Форматуємо числа до N знаків після коми
		precision := 2
		tmpl, _ := template.ParseFiles("templates/calculator2.html")
		tmpl.Execute(w, map[string]string{
			"mWnedA": roundFloat(mWnedA, precision),
			"mWnedP": roundFloat(mWnedP, precision),
			"mZper":  roundFloat(mZper, precision),
		})
		return
	}

	// Відображення HTML-сторінки
	tmpl, _ := template.ParseFiles("templates/calculator2.html")
	tmpl.Execute(w, nil)
}
