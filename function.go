package main

import "fmt"

func  greetings(name string) string  {
	return "Hello" + " " + name
}

func getSum(num1 int, num2 int)int  {
	return num1 + num1
}


func main()  {
	//fmt.Print(greetings("Obed"))
fmt.Println(getSum(3,6))
}
