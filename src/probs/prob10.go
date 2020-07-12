package probs

/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

//Seive of eratosthenes

func Prob10() uint64 {
	primes := []uint64{}

	var sum uint64 = 0

	primes = append(primes, 2)

	var i uint64 = 0
	for i = 3; i < 2000000; i = i + 2 {
		var isPrime bool = true

		for _, v := range primes {

			if v*v > i {
				break
			}

			if i%v == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, i)
		}
	}

	for _, v := range primes {
		sum = sum + v
	}

	return sum
}

/*
func generate(ch chan uint64) {
	var i uint64
	for i = 2; ; i++ {
		ch <- i
	}
}

func filter(in, out chan uint64, prime uint64) {
	for {
		i := <-in

		if i%prime != 0 {
			out <- i
		}
	}
}

func Prob10() uint64 {
	ch := make(chan uint64)
	var sum uint64 = 0

	go generate(ch)

	for {
		prime := <-ch

		//fmt.Print(prime, " ")
		if prime > 2000000 {
			break
		}
		sum = sum + prime

		ch1 := make(chan uint64)

		go filter(ch, ch1, prime)

		ch = ch1
		//time.Sleep(1e8)
	}

	return sum
}
*/
//Brute force::
/*func _isPrime(num int64) bool {

	var i int64 = 0

	for i = 2; i < num; i = i + 1 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func cal(num int64, limit int64, c chan int64) {

	var i int64 = 0
	var sum int64 = 0
	for i = num; i <= limit; i = i + 1 {
		if _isPrime(i) {
			sum = sum + i
		}
	}

	c <- sum
}

// 10000000 / 4 := 2500000

func Prob10() int64 {

	var sum int64 = 0
	ci := make(chan int64)

	//var num int64 = 0
	var a int64 = 0
	var b int64 = 0
	//var c int64 = 0
	//var d int64 = 0

	go cal(2, 5, ci)
	go cal(6, 10, ci)
	//go cal(2500000, ci)
	//go cal(2500000, ci)

	a, b = <-ci, <-ci
	sum = a + b
	return sum
}*/
