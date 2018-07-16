package strings

import "strconv"

type Fn func(a interface{}) string

func IntFn(a interface{}) string {
	return strconv.Itoa(a.(int))
}