package probs

import "fmt"

/*
A palindromic number reads the same both ways.
The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

*/

func _reverseNum(num int) int {
	var rev int = 0

	for num != 0 {
		rev = rev * 10
		rev = rev + (num % 10)
		num = num / 10
	}

	return rev
}

func _palindrome(num int) bool {

	rev := _reverseNum(num)

	if rev == num {
		return true
	} else {
		return false
	}

}

func Prob4() int {

	var high int = 0
	for i := 100; i <= 999; i++ {
		for j := i; j <= 999; j++ {
			if _palindrome(i * j) {
				if i*j > high {
					high = i * j
					fmt.Println(i, j)
				}
				//fmt.Println(i, j)
				//return i * j
			}
		}
	}
	return high
}
