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

// Функція інтегралу для обчислення первісної
func integral(a, b float64, n int, function func(float64) float64) float64 {
	h := (b - a) / float64(n)
	sum := 0.0
	for i := 0; i < n; i++ {
		x := a + float64(i)*h
		sum += function(x)
	}
	return sum * h
}

// Функція для обробки запиту на калькулятор 1
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Отримуємо значення з форми
		power, _ := strconv.ParseFloat(r.FormValue("Power"), 64)
		error1, _ := strconv.ParseFloat(r.FormValue("Error1"), 64)
		error2, _ := strconv.ParseFloat(r.FormValue("Error2"), 64)
		price, _ := strconv.ParseFloat(r.FormValue("Price"), 64)

		//задаємо значення границь інтегралу
		upperLimit := power + power*0.05
		lowerLimit := power - power*0.05

		//знаходимо значення інтегралу
		deltaW1 := integral(lowerLimit, upperLimit, 1000, func(x float64) float64 {
			return 1 / (error1 * math.Sqrt(2*math.Pi)) * math.Exp(-math.Pow(x-power, 2)/(2*math.Pow(error1, 2)))
		})

		//знаходимо значення прибутку та втрати до вдосконалення системи
		W1 := power * 24 * math.Round(deltaW1*100) / 100
		profit1 := W1 * price
		W2 := power * 24 * math.Round((1-deltaW1)*100) / 100
		lose1 := W2 * price

		result1 := profit1 - lose1

		deltaW2 := integral(lowerLimit, upperLimit, 1000, func(x float64) float64 {
			return 1 / (error2 * math.Sqrt(2*math.Pi)) * math.Exp(-math.Pow(x-power, 2)/(2*math.Pow(error2, 2)))
		})

		//знаходимо значення прибутку та втрати після вдосконалення системи
		W3 := power * 24 * math.Round(deltaW2*100) / 100
		profit2 := W3 * price
		W4 := power * 24 * math.Round((1-deltaW2)*100) / 100
		lose2 := W4 * price

		result2 := profit2 - lose2

		// Форматуємо числа до N знаків після коми
		precision := 2
		tmpl, _ := template.ParseFiles("templates/calculator.html")
		tmpl.Execute(w, map[string]string{
			"profit1": roundFloat(profit1, precision),
			"lose1":   roundFloat(lose1, precision),
			"result1": roundFloat(result1, precision),
			"profit2": roundFloat(profit2, precision),
			"lose2":   roundFloat(lose2, precision),
			"result2": roundFloat(result2, precision),
		})
		return
	}

	// Відображення HTML-сторінки
	tmpl, _ := template.ParseFiles("templates/calculator.html")
	tmpl.Execute(w, nil)
}
