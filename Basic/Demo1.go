package main

import "fmt"

func main() {
	const (
		a = iota * 2
		b
		c
		d
	)
	fmt.Println(a, b, c, d)

	const (
		A = 1
		_
		_
		_
		B = iota
		C
		D
		E = iota + 10
		F = 20
		G
		H = iota + 1
		M
	)
	fmt.Println(A, B, C, D, E, F, G, H, M)
}
