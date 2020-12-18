package main

import "fmt"

func main()  {
	for i := 0; i < 10; i++ {
		if  i %  2 ==0 {
			fmt.Println("Eveb", + i)
		}else if i % 3==0{
			fmt.Println("Odd", + i)
		}else {
			//fmt.Println(i)
			primeNumber(44)
		}
	}
}

func primeNumber( n int)  {
	for a:=11; a < n; a++ {
		fmt.Println("prime", + a)
	}
}