// combinatorial search
package main

import "fmt"

func main() {
	input := []int{1, 2, 3}
	a := make([]int, 3) // a is the partial solution
	backtrack(a, 0, input)
}

// backtrack generates subsets using Skiena's template
func backtrack(a []int, k int, input []int) {
	if isASolution(a, k, input) {
		processSolution(a, k, input)
	} else {
		// Two candidates at each position: include or exclude the next element
		c := make([]int, 2) // true=include, false=exclude
		constructCandidates(a, k, input, c)
		for _, ci := range c {
			a[k] = ci
			backtrack(a, k+1, input)
		}
	}
}

func isASolution(a []int, k int, input []int) bool {
	return k == len(input)
}

func processSolution(a []int, k int, input []int) {
	subset := []int{}
	for i := range k {
		if a[i] == 1 {
			subset = append(subset, input[i])
		}
	}
	fmt.Println(subset)
}

func constructCandidates(a []int, k int, input []int, c []int) {
	c[0] = 1  // include input[k]
	c[1] = 0  // exclude input[k]
}
