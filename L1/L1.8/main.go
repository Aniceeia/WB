package main

import "fmt"

func setBit(num int64, pos int, numberToSet int) (int64, error) {
	if pos < 0 || pos > 63 {
		return 0, fmt.Errorf("позиция бита должна быть от 0 до 63")
	}
	if numberToSet != 0 && numberToSet != 1 {
		return 0, fmt.Errorf("numberToSet должен быть 0 или 1")
	}
	mask := int64(1) << pos //создаем маску для установки бита в единицу
	if numberToSet == 1 {
		return num | mask, nil //используйем оператор или т.к. 1 or 0 = 1 и 1 or 1 = 1
	}
	return num &^ mask, nil //маска с единицей так, что для установки бита в 0 используем "не или" оператор 1 &^&1 = 0 и 1 &^ 0 = 0
}

func main() {
	result, err := setBit(5, 1, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	resultWithError1, err := setBit(5, -1, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultWithError1)

	resultWithError2, err := setBit(5, 1, 25)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultWithError2)

}
