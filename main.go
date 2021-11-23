package main

import "fmt"

type myNumbers []int

func (mn myNumbers) forEach(fn func(v int)) {
	for _, v := range mn {
		fn(v)
	}
}
func main() {
	var nums myNumbers = make([]int, 100)
	for i := 1; i <= 100; i++ {
		nums[i-1] = i
	}

	nums.forEach(func(v int) {

		if v%3 == 0 {
			fmt.Println("fizz")
			return
		}
		if v%5 == 0 {
			fmt.Println("buzz")
			return
		}

		fmt.Println(v)
	})
}
