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

// Функція для обробки запиту на калькулятор 1
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Отримуємо значення з форми
		H, _ := strconv.ParseFloat(r.FormValue("H"), 64)
		C, _ := strconv.ParseFloat(r.FormValue("C"), 64)
		S, _ := strconv.ParseFloat(r.FormValue("S"), 64)
		N, _ := strconv.ParseFloat(r.FormValue("N"), 64)
		O, _ := strconv.ParseFloat(r.FormValue("O"), 64)
		W, _ := strconv.ParseFloat(r.FormValue("W"), 64)
		A, _ := strconv.ParseFloat(r.FormValue("A"), 64)

		// Розрахунок коефіцієнтів переходу
		Kpc := 100 / (100 - W)
		Kpg := 100 / (100 - W - A)

		// Розрахунок нижчої теплоти згоряння для робочої маси
		Qrn := (339*C + 1030*H - 108.8*(O-S) - 25*W) / 1000

		// Розрахунок нижчої теплоти згоряння для сухої та горючої маси
		Qcn := (Qrn + 0.025*W) * Kpc
		Qgn := (Qrn + 0.025*W) * Kpg

		// Перерахунок складу сухої маси
		H_dry := H * Kpc
		C_dry := C * Kpc
		S_dry := S * Kpc
		N_dry := N * Kpc
		O_dry := O * Kpc
		A_dry := A * Kpc

		// Перерахунок складу горючої маси
		H_comb := H * Kpg
		C_comb := C * Kpg
		S_comb := S * Kpg
		N_comb := N * Kpg
		O_comb := O * Kpg

		// Форматуємо числа до N знаків після коми
		precision := 2
		precision2 := 4
		tmpl, _ := template.ParseFiles("templates/calculator.html")
		tmpl.Execute(w, map[string]string{
			"Kpc":    roundFloat(Kpc, precision),
			"Kpg":    roundFloat(Kpg, precision),
			"Qrn":    roundFloat(Qrn, precision2),
			"Qcn":    roundFloat(Qcn, precision2),
			"Qgn":    roundFloat(Qgn, precision2),
			"H_dry":  roundFloat(H_dry, precision),
			"C_dry":  roundFloat(C_dry, precision),
			"S_dry":  roundFloat(S_dry, precision),
			"N_dry":  roundFloat(N_dry, precision),
			"O_dry":  roundFloat(O_dry, precision),
			"A_dry":  roundFloat(A_dry, precision),
			"H_comb": roundFloat(H_comb, precision),
			"C_comb": roundFloat(C_comb, precision),
			"S_comb": roundFloat(S_comb, precision),
			"N_comb": roundFloat(N_comb, precision),
			"O_comb": roundFloat(O_comb, precision),
		})
		return
	}

	// Відображення HTML-сторінки
	tmpl, _ := template.ParseFiles("templates/calculator.html")
	tmpl.Execute(w, nil)
}
