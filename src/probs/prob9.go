package probs

/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a*a + b*b = c*c
For example, 3*3 + 4*4 = 9 + 16 = 25 = 5*5.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

func Prob9() int {

	sum := 1000
	a := 1
	b := 1
	c := 0
	found := false

	for a = 1; a < sum/3; a = a + 1 {

		for b = a; b < sum/2; b = b + 1 {

			c = sum - a - b

			if c*c == a*a+b*b {
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	return a * b * c

}
