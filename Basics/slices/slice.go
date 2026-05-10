package main

import (
	f "fmt"
	"slices"
)

func main(){
	var s  []string
	f.Println("Uninitiated slice", s, len(s), s==nil) 

	a := make([]string, 4)
	f.Println("Initiated slice", a, len(a), cap(a))

	a[0] = "Hello"
	a[1] = "There"
	f.Println(a)

	a = append(a, "My name")
	a = append(a, "is Devansh")
	f.Println(a, len(a))

	c := make([] string, len(a))
	copy(c,a)
	f.Println(c)

	l := c[:2]
	f.Println(l)
	x := c[3:6]
	f.Println(x)
	y := c[2:]
	f.Println(y)
    z := c[2:]
	f.Println(slices.Equal(y, z))

	var twoD = make([][]int, 4)
	for i := range len(twoD){
		innerLen := i+1
		twoD[i] = make([] int, innerLen)
		for j:= range innerLen {
			twoD[i][j] = i+j
		}
	}
	f.Println(twoD)
}