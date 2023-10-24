package main

import "fmt"

type group []int

func (g group) checkEvenOrOdd() {
	for _, v := range g {
		// intGroup[i] = i
		if v%2 == 0 {
			fmt.Printf("%v is even\n", v)
		} else {
			fmt.Printf("%v is odd\n", v)
		}
	}

}

func main() {

	intGroup := make(group, 11)

	for i := range intGroup {
		intGroup[i] = i
	}
	fmt.Println(intGroup)
	intGroup.checkEvenOrOdd()

}
