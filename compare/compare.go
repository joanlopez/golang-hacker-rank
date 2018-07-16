package compare

type Fn func(a, b interface{}) int

func IntFn(a, b interface{}) int {
	if a.(int) > b.(int) {
		return 1
	} else if a.(int) < b.(int) {
		return -1
	}
	return 0
}

func StringFn(a, b interface{}) int {
	if a.(string) > b.(string) {
		return 1
	} else if a.(string) < b.(string) {
		return -1
	}
	return 0
}