package main

import "fmt"



func main(){

	// Basic Loop
	i := 0
 for i<=3 {
	fmt.Println(i)
	i++
}
fmt.Println(i)
// Conditional loop
for j:=0; j<3; j++{
	fmt.Println(j)
}

// Range loop
for i := range 4 {
	fmt.Println("Range : ", i)
}

// Continuation and break
for n := range 6 {
	if n%2 == 0 {
		continue
	}
	fmt.Println(n)
}

}

