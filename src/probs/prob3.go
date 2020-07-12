package probs

import "math"

/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/

func Prob3() int64 {

	var n int64 = 600851475143
	var high int64 = 0

	var i int64 = 0

	for n%2 == 0 {

		high = 2
		n = n / 2
	}

	var x float64 = float64(n)

	var b = int64(math.Sqrt(x))

	for i = 3; i < b; i = i + 2 {

		for n%i == 0 {
			if i > high {
				high = i
			}
			n = n / i
		}
	}

	return high
}
