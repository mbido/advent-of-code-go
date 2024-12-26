package utils

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/hmdsefi/gograph"
)

// Example of a generic function
func PrintHello() {
	fmt.Println("Hello, Advent of Code!")
}

func Nums(s string) []int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(s, -1)
	var result []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}

func SNums(s string) []int {
	re := regexp.MustCompile(`[-+]?\d+`)
	matches := re.FindAllString(s, -1)
	var result []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}

func Floats(s string) []float64 {
	re := regexp.MustCompile(`[-+]?\d*\.?\d+`)
	matches := re.FindAllString(s, -1)
	var result []float64
	for _, match := range matches {
		num, err := strconv.ParseFloat(match, 64)
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func CountInt(l []int, target int) int {
	count := 0
	for _, n := range l {
		if n == target {
			count++
		}
	}
	return count
}

func IsSorted[T any](l []T, f func(T, T) bool) bool {
	for i := 1; i < len(l); i++ {
		if !f(l[i-1], l[i]) {
			return false
		}
	}
	return true
}

func RemoveAtIndex[T any](l *[]T, index int) {
	if index < 0 || index >= len(*l) {
		return
	}
	*l = append((*l)[:index], (*l)[index+1:]...)
}

func AddAtIndex[T any](l *[]T, index int, value T) {
	if index < 0 || index > len(*l) {
		return
	}
	*l = append((*l)[:index], append([]T{value}, (*l)[index:]...)...)
}

func Sum(l []int) int {
	sum := 0
	for _, n := range l {
		sum += n
	}
	return sum
}

func main() {
	input := "Hello, Advent of Code 2021! Here are some numbers: 123, 456, and 789."
	fmt.Println(Nums(input)) // Output: [2021, 123, 456, 789]
	graph := gograph.New[int](gograph.Directed())
	graph.AddEdge(gograph.NewVertex(1), gograph.NewVertex(2))
}
