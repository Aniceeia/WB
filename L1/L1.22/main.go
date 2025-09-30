package main

import (
	"fmt"
	"math/big"
	"strings"
)

// чтобы результат операций выглядил читабельно преобразуем получившееся большое число в экспоненциальную форму записи
func toExponentialPrecise(num *big.Int, precision int) string {
	if num.Sign() == 0 {
		return "0"
	}

	str := num.String()
	if len(str) <= 5 { //если число меньше 100000, то возвращаем его в обычной форме записи
		return str
	}
	//делим строку на числа до запятой и после
	firstDigit := string(str[0])
	remaining := str[1:]
	//степень экспоненты - длина строки без первого элемента
	exponent := len(str) - 1

	if precision > 0 && len(remaining) > precision { // в зависимости от выбранной нами точности записываем числа в правую часть числа
		remaining = remaining[:precision]
	}

	trimmedRemaining := strings.TrimRight(remaining, "0") //удаляем оставшиеся числа
	if trimmedRemaining == "" {
		return fmt.Sprintf("%se%d", firstDigit, exponent)
	}

	return fmt.Sprintf("%s.%se%d", firstDigit, trimmedRemaining, exponent)
}

func main() {
	a := new(big.Int)
	b := new(big.Int)

	two := big.NewInt(2)
	exponentA := big.NewInt(70)
	exponentB := big.NewInt(53)

	a.Exp(two, exponentA, nil)
	b.Exp(two, exponentB, nil)

	result := new(big.Int)

	fmt.Println("a   =", toExponentialPrecise(a, 6))
	fmt.Println("b   =", toExponentialPrecise(b, 6))
	fmt.Println("add =", toExponentialPrecise(result.Add(a, b), 6))
	fmt.Println("sub =", toExponentialPrecise(result.Sub(a, b), 6))
	fmt.Println("div =", toExponentialPrecise(result.Div(a, b), 6))
	fmt.Println("mul =", toExponentialPrecise(result.Mul(a, b), 6))
}
