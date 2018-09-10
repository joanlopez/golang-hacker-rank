package strings

import (
	"strconv"
	"strings"
)

type Fn func(a interface{}) string

func IntFn(a interface{}) string {
	return strconv.Itoa(a.(int))
}

func JoinByIgnoring(s []string, by, ignore string) string {
	var r []string
	for _, str := range s {
		if str != ignore {
			r = append(r, str)
		}
	}
	return strings.Join(r, by)
}