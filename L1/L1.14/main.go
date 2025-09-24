package main

import "fmt"

func typeChecker(i any) {
	switch i.(type) {
	case int:
		fmt.Printf("%-15v int\n", i)
	case string:
		fmt.Printf("%-15v string\n", i)
	case bool:
		fmt.Printf("%-15v bool\n", i)
	case chan int:
		fmt.Printf("%-15v chan int\n", i)
	default:
		fmt.Println("unknown type")
	}
}

func main() {
	intType := 1
	stringType := "string"
	boolType := true
	chanType := make(chan int)

	typeChecker(intType)
	typeChecker(stringType)
	typeChecker(boolType)
	typeChecker(chanType)

}
