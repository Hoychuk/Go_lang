package handlers

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)

// Функція для округлення чисел до N знаків після коми
func roundFloat(value float64, precision int) string {
	format := fmt.Sprintf("%%.%df", precision)
	return fmt.Sprintf(format, value)
}

// Функція для обробки запиту на калькулятор 1
func Task1Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		U := 10.0

		// Отримуємо значення з форми
		I, _ := strconv.ParseFloat(r.FormValue("I"), 64)
		t, _ := strconv.ParseFloat(r.FormValue("t"), 64)
		Sm, _ := strconv.ParseFloat(r.FormValue("Sm"), 64)
		jEk, _ := strconv.ParseFloat(r.FormValue("jEk"), 64)

		//Проводимо розрахунки згідно формул
		Im := Sm / 2 / (math.Sqrt(3.0) * U)
		ImPA := 2 * Im
		sEk := Im / jEk
		s := I * 1000 * math.Sqrt(t) / 92

		// Форматуємо числа до N знаків після коми
		precision := 2
		tmpl, _ := template.ParseFiles("templates/calculator1.html")
		tmpl.Execute(w, map[string]string{
			"Im":   roundFloat(Im, precision),
			"ImPA": roundFloat(ImPA, precision),
			"sEk":  roundFloat(sEk, precision),
			"s":    roundFloat(s, precision),
		})
		return
	}

	// Відображення HTML-сторінки
	tmpl, _ := template.ParseFiles("templates/calculator1.html")
	tmpl.Execute(w, nil)
}
