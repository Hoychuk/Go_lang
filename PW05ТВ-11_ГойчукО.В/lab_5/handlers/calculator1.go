package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Функція для округлення чисел до N знаків після коми
func roundFloat(value float64, precision int) string {
	format := fmt.Sprintf("%%.%df", precision)
	return fmt.Sprintf(format, value)
}

// Структура параметрів надійності
type ReliabilityParameters struct {
	omega float64
	tV    int
	mu    float64
	tP    int
}

// Дані для розрахунків
var data = map[string]ReliabilityParameters{
	"PL-110 kV":                    {0.007, 10, 0.167, 35},
	"PL-35 kV":                     {0.02, 8, 0.167, 35},
	"PL-10 kV":                     {0.02, 10, 0.167, 35},
	"CL-10 kV (Trench)":            {0.03, 44, 1.0, 9},
	"CL-10 kV (Cable Channel)":     {0.005, 18, 1.0, 9},
	"T-110 kV":                     {0.015, 100, 1.0, 43},
	"T-35 kV":                      {0.02, 80, 1.0, 28},
	"T-10 kV (Cable Network)":      {0.005, 60, 0.5, 10},
	"T-10 kV (Overhead Network)":   {0.05, 60, 0.5, 10},
	"B-110 kV (Gas-Insulated)":     {0.01, 30, 0.1, 30},
	"B-10 kV (Oil)":                {0.02, 15, 0.33, 15},
	"B-10 kV (Vacuum)":             {0.05, 15, 0.33, 15},
	"Busbars 10 kV per Connection": {0.03, 2, 0.33, 15},
	"AV-0.38 kV":                   {0.05, 20, 1.0, 15},
	"ED 6,10 kV":                   {0.1, 50, 0.5, 0},
	"ED 0.38 kV":                   {0.1, 50, 0.5, 0},
}

// Функція для обробки запиту на калькулятор 1
func TaskHandler1(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Отримуємо значення з форми

		wOc := 0.0
		tVOc := 0.0

		for key := range data {
			amount, _ := strconv.Atoi(r.FormValue(key))
			indicator := data[key]

			if amount > 0 {
				wOc += float64(amount) * indicator.omega
				tVOc += float64(amount) * indicator.omega * float64(indicator.tV)
			}
		}

		//Обчислення
		tVOc /= wOc
		kAOc := (tVOc * wOc) / 8760
		kPOs := 1.2 * 43 / 8760
		wDk := 2 * wOc * (kAOc + kPOs)
		wDc := wDk + 0.02

		// Форматуємо числа до N знаків після коми
		precision := 5
		tmpl, _ := template.ParseFiles("templates/calculator1.html")
		tmpl.Execute(w, map[string]string{
			"wOc":  roundFloat(wOc, precision),
			"tVOc": roundFloat(tVOc, precision),
			"kAOc": roundFloat(kAOc, precision),
			"kPOs": roundFloat(kPOs, precision),
			"wDk":  roundFloat(wDk, precision),
			"wDc":  roundFloat(wDc, precision),
		})
		return
	}

	// Відображення HTML-сторінки
	tmpl, _ := template.ParseFiles("templates/calculator1.html")
	tmpl.Execute(w, nil)
}
