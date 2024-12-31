package main

import "fmt"

func q1() {
	var i int = 20
	var f float64 = float64(i)
	fmt.Println(i, f)
}

func q2() {
	const value = 10
	var i int = value
	var f float64 = value
	fmt.Println(i, f)
}

func q3() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	b += 1
	smallI += 1
	bigI += 1
	fmt.Println(b, smallI, bigI)
}

func main() {
	q1()
	q2()
	q3()
}
