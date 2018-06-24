package main

import (
	"fmt"
	"strings"
)

func Staircase(n int32) {
	var height = int(n)
	for i := 1; i <= height; i++ {
		fmt.Printf(
			"%s%s\n",
			strings.Repeat(" ", height-i),
			strings.Repeat("#", i),
		)
	}
}
