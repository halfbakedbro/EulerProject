package probs

/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

func isPrime(can int32) bool {

	var i int32 = 0

	for i = 2; i < can; i = i + 1 {
		if can%i == 0 {
			return false
		}
	}
	return true
}

func Prob7() int32 {

	var can int32 = 0
	var count int32 = 0

	for can, count = 2, 0; count < 10001; can = can + 1 {
		if isPrime(can) {
			count = count + 1
		}
	}

	return can - 1
}
