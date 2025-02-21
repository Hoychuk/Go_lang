package handlers

import (
	"html/template"
	"math"
	"net/http"
	"strconv"
)

// Функція для обробки запиту на калькулятор 3
func Task3Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		SNom := 6.3
		Uh := 115.0
		Ul := 11.0
		UkMax := 11.1

		// Отримуємо значення з форми
		Rcn, _ := strconv.ParseFloat(r.FormValue("Rcn"), 64)
		Xcn, _ := strconv.ParseFloat(r.FormValue("Xcn"), 64)
		RcMin, _ := strconv.ParseFloat(r.FormValue("RcMin"), 64)
		XcMin, _ := strconv.ParseFloat(r.FormValue("XcMin"), 64)

		//Проводимо розрахунки згідно формул
		XT := UkMax * math.Pow(Uh, 2.0) / 100 / SNom
		Xsh := Xcn + XT
		Zsh := math.Sqrt(math.Pow(Rcn, 2.0) + math.Pow(Xsh, 2.0))

		XshMin := XcMin + XT
		ZshMin := math.Sqrt(math.Pow(RcMin, 2.0) + math.Pow(XshMin, 2.0))

		Ish3 := Uh * 1000 / math.Sqrt(3.0) / Zsh
		Ish2 := Ish3 * math.Sqrt(3.0) / 2

		IshMin3 := Uh * 1000 / math.Sqrt(3.0) / ZshMin
		IshMin2 := IshMin3 * math.Sqrt(3.0) / 2

		Kpr := math.Pow(Ul, 2.0) / math.Pow(Uh, 2.0)

		RshN := Rcn * Kpr
		XshN := Xsh * Kpr
		ZshN := math.Sqrt(math.Pow(RshN, 2.0) + math.Pow(XshN, 2.0))

		RshMinN := RcMin * Kpr
		XshMinN := XshMin * Kpr
		ZshMinN := math.Sqrt(math.Pow(RshMinN, 2.0) + math.Pow(XshMinN, 2.0))

		IshN3 := Ul * 1000 / math.Sqrt(3.0) / ZshN
		IshN2 := IshN3 * math.Sqrt(3.0) / 2

		IshMinN3 := Ul * 1000 / math.Sqrt(3.0) / ZshMinN
		IshMinN2 := IshMinN3 * math.Sqrt(3.0) / 2

		Rl := 7.91
		Xl := 4.49

		RSumN := Rl + RshN
		XSumN := Xl + XshN
		ZSumN := math.Sqrt(math.Pow(RSumN, 2.0) + math.Pow(XSumN, 2.0))

		RSumMinN := Rl + RshMinN
		XSumMinN := Xl + XshMinN
		ZSumMinN := math.Sqrt(math.Pow(RSumMinN, 2.0) + math.Pow(XSumMinN, 2.0))

		Iln3 := Ul * 1000 / math.Sqrt(3.0) / ZSumN
		Iln2 := Iln3 * math.Sqrt(3.0) / 2

		IlMinN3 := Ul * 1000 / math.Sqrt(3.0) / ZSumMinN
		IlMinN2 := IlMinN3 * math.Sqrt(3.0) / 2

		// Форматуємо числа до N знаків після коми
		precision := 2
		tmpl, _ := template.ParseFiles("templates/calculator3.html")
		tmpl.Execute(w, map[string]string{
			"Zsh":      roundFloat(Zsh, precision),
			"ZshMin":   roundFloat(ZshMin, precision),
			"Ish3":     roundFloat(Ish3, precision),
			"Ish2":     roundFloat(Ish2, precision),
			"IshMin3":  roundFloat(IshMin3, precision),
			"IshMin2":  roundFloat(IshMin2, precision),
			"Kpr":      roundFloat(Kpr, precision),
			"ZshN":     roundFloat(ZshN, precision),
			"ZshMinN":  roundFloat(ZshMinN, precision),
			"IshN3":    roundFloat(IshN3, precision),
			"IshN2":    roundFloat(IshN2, precision),
			"IshMinN3": roundFloat(IshMinN3, precision),
			"IshMinN2": roundFloat(IshMinN2, precision),
			"ZSumN":    roundFloat(ZSumN, precision),
			"ZSumMinN": roundFloat(ZSumMinN, precision),
			"Iln3":     roundFloat(Iln3, precision),
			"Iln2":     roundFloat(Iln2, precision),
			"IlMinN3":  roundFloat(IlMinN3, precision),
			"IlMinN2":  roundFloat(IlMinN2, precision),
		})
		return
	}

	// Відображення HTML-сторінки
	tmpl, _ := template.ParseFiles("templates/calculator3.html")
	tmpl.Execute(w, nil)
}
