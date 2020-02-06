package main

func add(a int, b int) int {
	return a + b
}

func sumList(l []int) int {
	answer := 0
	for a, _ := range l {
		answer += a
	}
	return answer
}
