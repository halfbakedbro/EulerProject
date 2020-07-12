package probs

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

func Prob5() int64 {

	var i int64 = 0
	for i = 1; i <= 9999999999; i++ {

		if i%16 == 0 {
			// then divisible by 2 4 8 16

			if i%18 == 0 {
				// then divisible by 3 9 18

				if i%15 == 0 {
					// then divisible by 5 15
					if i%6 == 0 && i%7 == 0 && i%10 == 0 && i%11 == 0 {

						if i%12 == 0 && i%13 == 0 && i%14 == 0 && i%17 == 0 && i%19 == 0 && i%20 == 0 {
							return i
						}
					}
				}
			}

		}

	}

	return -1
}
