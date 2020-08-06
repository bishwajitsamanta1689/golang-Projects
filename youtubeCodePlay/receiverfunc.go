package main

import "fmt"

type person struct {
	first string
	last string
	saying string
}
// Receiver Function Syntax
// func (Receiver) identifier(Parameters) method() return output {}

func (p person) speak() string{
	return p.first + " says " + p.saying
}

func main(){
	p1:= person{
		first: "Bishwajit",
		last: "samanta",
		saying: "hey buddy!!",
	}
	p2:= person{
		first:  "Payel",
		last:   "Samanta",
		saying: "Come Here!!",
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("================")
	fmt.Println(p1.speak())
	fmt.Println(p2.speak())

}


