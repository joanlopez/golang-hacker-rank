package implementation

func Kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v1 > v2 && ((x2-x1)%(v1-v2) == 0) {
		return "YES"
	}
	return "NO"
}
