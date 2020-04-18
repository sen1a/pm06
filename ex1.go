package main

import (
  "fmt"
  "math"
)

const course float64 = 92.0456

func main() {
	fmt.Print("Введите сумму в рублях:")
  var money float64
  fmt.Scanf("%f", &money)
  money /= course
  money = money * 100
  money = math.Round(money)
  money = money / 100
  fmt.Printf("Вы получите %.2f Фунтов стерлингов Соединенного королевства\n", money);
}
