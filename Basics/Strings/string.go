package main

import "fmt"
func main(){
	s := "hELLO"
	fmt.Println("length :", len(s))

	for i:=0; i<len(s); i++ {
		fmt.Printf("%d : %x\n", i, s[i])
	}

	x := 'X'
	fmt.Printf("x :", x)
}
