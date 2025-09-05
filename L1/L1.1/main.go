package main

import "fmt"

type Human struct {
	Name       string
	Age        int
	ZodiacSign string
	WorkPlace  string
}

func (h *Human) SayHello() {
	fmt.Printf("hello, I'm %s", h.Name)
}

func (h *Human) SaySign() {
	fmt.Printf("I'm %s", h.ZodiacSign)
}

func (h *Human) SayAge() {
	fmt.Printf("I'm %d y.o.", h.Age)
}

func (h *Human) SayWorkPlace() {
	fmt.Printf("I'm currently working in %s", h.WorkPlace)
}

type Action struct {
	Human
	KindOfAction string
}

func (a *Action) ReviewHuman() {
	fmt.Printf("Name: %s, Age: %d, Sign: %s, Workplace: %s\n", a.Name, a.Age, a.ZodiacSign, a.WorkPlace)
}

func main() {
	action := Action{
		Human: Human{
			Name:       "Arthur",
			Age:        27,
			ZodiacSign: "Libra",
			WorkPlace:  "WB",
		},
		KindOfAction: "learning",
	}

	fmt.Printf("Use of parent funcs in main\nName: %s, Age: %d, Sign: %s, Workplace: %s\n", action.Name, action.Age, action.ZodiacSign, action.WorkPlace)
	fmt.Printf("Use of parent funcs in action funcs\n")
	action.ReviewHuman()
	fmt.Printf("Use of parent funcs mixed with Acrion structs elements\n")
	fmt.Printf("So, your name is %s, age %d, sign %s and workplase is %s? Fine, but you should %s to be better progger\n", action.Name, action.Age, action.ZodiacSign, action.WorkPlace, action.KindOfAction)
}
