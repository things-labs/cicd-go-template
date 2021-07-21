package cpt

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func Add2(a, b int) int {
	a = a + b
	return a
}
