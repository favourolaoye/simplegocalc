package grader

func Grader(score float32, matric string) string{
	if score >= 70 && score <= 100{
		return "A"
	}else if score >= 60{
		return "B"
	}else if score >= 50 {
		return "C"
	}else if score >= 40 {
		return "D"
	}else {
		return "Failed"
	}
}