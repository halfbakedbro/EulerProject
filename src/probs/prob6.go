package probs

/*

The sum of the squares of the first ten natural numbers is,

12+22+...+102=385
The square of the sum of the first ten natural numbers is,

(1+2+...+10)2=552=3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025−385=2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

//sum of n natural numbers : n(n+1)/2

// sum of square of n natural numbers : ((2n+1)n(n+1))/6

func Prob6() int32 {

	var sum int32 = 0
	var sum_n int32 = 0

	sum = 100 * (100 + 1) / 2
	sum = sum * sum

	sum_n = ((2*100 + 1) * 100 * (100 + 1)) / 6

	return sum - sum_n
}
