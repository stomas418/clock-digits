package main

import (
	"sort"
)

func alreadyUsed(u [][]int, n []int) bool {
	similarities := 0
	for i := 0; i < len(u); i++ {
		similarities = 0
		for j := 0; j < 4; j++ {
			if u[i][j] == n[j] {
				similarities += 1
			}
		}
		if similarities == 4 {
			return true
		}
	}
	return false
}

//The solution works looping through the arguments, trying each digit in each position of a digital clock time format.
func solution(A, B, C, D int) int {

	arr := []int{A, B, C, D} //i organize all arguments into an array for ease to loop
	var used [][]int         //this var will contain all used solutions so we don't end up repeating them
	sort.Ints(arr)           //order A, B, C, D from lowest to highest
	var result int
	for i := 0; i < 4; i++ {
		first := arr[i]
		if first >= 3 { //if the lowest number in the arr is bigger than 3 then no other number is going to be suitable for the clock
			return result
		}
		for j := 0; j < 4; j++ {
			if j == i {
				continue
			}
			second := arr[j]
			if first == 2 && second > 3 {
				continue
			}
			for k := 0; k < 4; k++ {
				if k == j || k == i {
					continue
				}
				third := arr[k]
				if third > 5 {
					continue
				}
				for l := 0; l < 4; l++ {
					if l == i || l == j || l == k { //if i, j, k, or l are the same, we are comparing at least one number with itself, so it's not a valid operation
						continue
					}
					fourth := arr[l]
					if fourth > 9 {
						return 0
					}
					resArr := []int{first, second, third, fourth}
					if !alreadyUsed(used, resArr) {
						result += 1
						used = append(used, resArr)
					}
				}
			}
		}
	}
	return result
}
