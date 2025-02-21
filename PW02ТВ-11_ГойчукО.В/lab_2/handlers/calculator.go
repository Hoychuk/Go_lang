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
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Отримуємо значення з форми
		coal, _ := strconv.ParseFloat(r.FormValue("Coal"), 64)
		masut, _ := strconv.ParseFloat(r.FormValue("Masut"), 64)
		gas, _ := strconv.ParseFloat(r.FormValue("Gas"), 64)

		// Розраховуємо коефіцієнти емісії та валовий викид твердих частинок вугілля
		kCoal := math.Pow(10, 6) / 20.47 * 0.8 * 25.2 / (100 - 1.5) * (1 - 0.985)
		ECoal := math.Pow(10, -6) * kCoal * 20.47 * coal
		// --\\-- мазуту
		kMasut := math.Pow(10, 6) / 39.48 * 1 * 0.15 / (100 - 0) * (1 - 0.985)
		EMasut := math.Pow(10, -6) * kMasut * 39.48 * masut
		// --\\-- природнього газу
		kGas := math.Pow(10, 6) / 33.08 * 0 * 0 / (100 - 0) * (1 - 0.985)
		EGas := math.Pow(10, -6) * kGas * 33.08 * gas

		// Форматуємо числа до N знаків після коми
		precision := 2
		tmpl, _ := template.ParseFiles("templates/calculator.html")
		tmpl.Execute(w, map[string]string{
			"kCoal":  roundFloat(kCoal, precision),
			"ECoal":  roundFloat(ECoal, precision),
			"kMasut": roundFloat(kMasut, precision),
			"EMasut": roundFloat(EMasut, precision),
			"kGas":   roundFloat(kGas, precision),
			"EGas":   roundFloat(EGas, precision),
		})
		return
	}

	// Відображення HTML-сторінки
	tmpl, _ := template.ParseFiles("templates/calculator.html")
	tmpl.Execute(w, nil)
}
