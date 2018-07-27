package implementation

func nextMultipleOf(n int32, multipleOf int32) int32 {
	if n % multipleOf == 0 {
		return n
	}
	return n + (5 - n % 5)
}

// TODO: Improve with a rule validator
func roundGrade(grade int32) int32 {
	nextMultiple := nextMultipleOf(grade, 5)
	if nextMultiple - grade > 2 || grade < 38 {
		return grade
	}
	return nextMultiple
}

func GradingStudents(grades []int32) []int32 {
	for i := range grades {
		grades[i] = roundGrade(grades[i])
	}
	return grades
}
