package main

import (
	"fmt"
	"math/big"
)

// Для работы с очень большими числами в го, есть пакет big
func main() {
	var sign string
	a := new(big.Int)
	b := new(big.Int)
	zero := big.NewInt(0)
	res := new(big.Int)

	fmt.Println("please, enter the first big number")
	fmt.Scanln(a)
	fmt.Println("please, enter the second big number")
	fmt.Scanln(b)
	fmt.Println("please, enter the arithmetic sign")
	fmt.Scanln(&sign)

	if sign == "-" {
		b = b.Neg(b)
		res := res.Add(a, b)
		fmt.Printf("%d%d=%d\n", a, b, res)
	} else if sign == "+" {
		res := res.Add(a, b)
		fmt.Printf("%d+%d=%d\n", a, b, res)
	} else if sign == "*" {
		res := res.Mul(a, b)
		fmt.Printf("%d*%d=%d\n", a, b, res)
	} else if sign == "/" {
		if b.Cmp(zero) == 0 {
			fmt.Printf("you shouldn't divide by zero!")
			return
		}
		res := res.Div(a, b)
		fmt.Printf("%d/%d=%d\n", a, b, res)
	}
}
