package main
import (
	"fmt"
)

// 1.1.14
func histogram(a []int, m int) []int {
	result := make([]int, m)
	fmt.Pri
	for i := 0; i < len(a); i++ {
		result[a[i]] = result[a[i]] + 1;
	}
	return result
}

func test_histogram() {
	fmt.Println("Excercise 1.1.14")
	input := []int{3, 3, 2, 3, 1, 1, 1, 3, 4}
	m := 5
	fmt.Println(input)
	result := histogram(input, m)
	fmt.Println(result)
}

func main() {
	test_histogram()
}
