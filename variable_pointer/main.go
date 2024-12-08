package main

import (
	"fmt"
)

const secret = "abc"

type Os int

const (
	Mac Os = iota + 1
	Windows
	Linux
)

var (
	ii int
	ss string
	bb bool
)

func main() {
	var i int
	fmt.Printf("i: %v %T\n", i, i)

	var j int = 2
	fmt.Printf("j: %v %T\n", j, j)

	var k = 3
	fmt.Printf("k: %v %T\n", k, k)

	l := 4
	fmt.Printf("l: %v %T\n", l, l)

	ui := uint16(2)
	fmt.Printf("ui: %v %T\n", ui, ui)

	f := 1.23456
	s := "hello"
	b := true
	fmt.Printf("f: %[1]v %[1]T\n", f)
	fmt.Printf("s: %[1]v %[1]T\n", s)
	fmt.Printf("b: %[1]v %[1]T\n", b)

	pi, title := 3.14, "Go"
	fmt.Printf("pi: %[1]v title: %[2]v\n", pi, title)

	x := 10
	y := 1.23
	z := float64(x) + y
	fmt.Printf("z: %v\n", z)

	fmt.Printf("Mac: %[1]v Windows: %[2]v Linux: %[3]v\n", Mac, Windows, Linux)

	i = 2
	fmt.Printf("i: %v %T\n", i, i)
	i += 1
	fmt.Printf("i: %v %T\n", i, i)
	i *= 2
	fmt.Printf("i: %v %T\n", i, i)
}
