package main

import "fmt"

func main() {

	var x float32
	var y float32
	x, y = 2.0, 2.5
	var c3 = complex(x, y)
	fmt.Printf("C3 Type %T", c3)
	fmt.Println("c3", c3)
	r := 5.1
	i := 7.2
	c1 := complex(r, i)
	c2 := 6.1 + 2.1i
	fmt.Println("Hello c1:", c1, "C2: ", c2)

	cadd := c1 + c2
	csub := c1 - c2
	cmul := c1 * c2
	cdiv := c1 / c2
	fmt.Println("cadd", cadd)
	fmt.Println("csub", csub)
	fmt.Println("cmul", cmul)
	fmt.Println("cdiv", cdiv)
}
