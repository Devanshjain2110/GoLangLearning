package main
import "fmt"

func intSeq() func() int {
	i := 0;
	return func() int {
		i++
		return i
	}
	}
func main() {
	 intNext := intSeq()
	 fmt.Println(intNext()) 
	 fmt.Println(intNext()) 
	 fmt.Println(intNext()) 

	 intNexts := intSeq()
fmt.Println(intNexts())
	}

