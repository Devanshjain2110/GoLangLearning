
package main
import "fmt"
func main(){

	var a[5] int
	fmt.Println("Initial Array :", a)

	a[4] = 53
	fmt.Println("After setting a value a will be: ", a)

	Alen := len(a)
	fmt.Println("The length of array is ", Alen)

	b := [5] int{1,2,3,4,5}
	fmt.Println("The directly initialized with values array is", b)

	c := [...] int{32,4,3324,42,1313,1313,13,2}
	fmt.Println("The compiled type initialized array looks like :", c, "and the length of that array is ", len(c))

	var twoD [2][3]int
	for i:= range 2 {
		for j := range 3 {
			twoD[i][j] = i+j
		}
	}

	fmt.Println("Our intiialized two d array looks like this ", len(twoD[0]))

}