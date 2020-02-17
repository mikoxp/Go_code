package function

func Add(a int,b int) int{
	return a+b
}

func Factor(n int) int {
	if n<=1{
		return 1
	}
	return Factor(n-1)*n
}