package implementation

import "fmt"

func in(n, x, y int32) bool {
	return n >= x && n <= y
}

func CountFruit(start, end, tree int32, fruits []int32) int32 {
	var count int32 = 0
	for _, fruit := range fruits {
		if in(tree+fruit, start, end) {
			count += 1
		}
	}
	return count
}

func CountApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	fmt.Printf("%v\n", CountFruit(s, t, a, apples))
	fmt.Printf("%v\n", CountFruit(s, t, b, oranges))
}