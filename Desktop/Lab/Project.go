package main

import (
	"fmt"
)

var R1 [5]int
var R2 [5]int
var R3 [5]int

var R1b [2]int
var R2b [2]int
var R3b [2]int

var l0 [3]int
var l1 [3]int
var l2 [3]int
var l3 [3]int
var l4 [3]int

func main() {
	//location := [5]int{0, 1, 2, 3, 4}
	
	R1 = [5]int{5, 3, 2, 7, 1}
	R2 = [5]int{4, 2, 5, 6, 3}
	R3 = [5]int{3, 5, 8, 1, 4}

	l0 = [3]int{5, 4, 3}
	l1 = [3]int{3, 2, 5}
	l2 = [3]int{2, 5, 8}
	l3 = [3]int{7, 6, 1}
	l4 = [3]int{1, 3, 4}
	LOC := 0
	min_d, rep := min_dist(LOC)
	isavail := busy(rep)
	for rep != -1 {
		if isavail == true {
			fmt.Printf("Representative %d is assigned\n", rep)
		} else {
			min_d, rep = min_dist(LOC)
		}
	}
	if rep == -1 {
		fmt.Println("No Representtives available\n")
	}

}

func busy(rep_num int) bool {
	if rep_num == 1 {
		if R1b[0] == 1 {
			return false
		} else {
			return true
		}
	} else if rep_num == 2 {
		if R2b[0] == 1 {
			return false
		} else {
			return true
		}
	} else if rep_num == 3 {
		if R3b[0] == 1 {
			return false
		} else {
			return true
		}
	}
	return false
}

func min_dist(x int) (int, int) {
	distance := -1
	rep_num := -1
	if x == 0 && busy(0) {
		distance, rep_num = min(l0)
	} else if x == 1 && busy(1) {
		distance, rep_num = min(l1)
	} else if x == 2 && busy(2) {
		distance, rep_num = min(l2)
	} else if x == 3 && busy(3) {
		distance, rep_num = min(l3)
	} else if x == 4 && busy(4) {
		distance, rep_num = min(l4)
	}
	return distance, rep_num
}

// func min_distance(x int) (int, string) {
// 	if (R1[x]) < R2[x] && R1[x] < R3[x] {
// 		return R1[x], "R1"
// 	} else if R2[x] < R1[x] && R2[x] < R3[x] {
// 		return R2[x], "R2"
// 	} else {
// 		return R3[x], "R3"
// 	}
// }

func min(x [3]int) (int, int) {
	min := x[0]
	temp := 0
	for i := 1; i < 3; i++ {
		if x[i] < min {
			min = x[i]
			temp = i
		}
	}
	return min, temp + 1
}

// fmt.Printf("%d %s", min_d, rep)
// if rep == 1 {
// 	if R1b[0] == 0 {
// 		R1b[0] = 1
// 		fmt.Println("Rep 1 is assigned\n")
// 	} else if rep == 2 {
// 		if R2b[0] == 0 {
// 			R2b[0] = 1

// 		}
// 	}
// }
