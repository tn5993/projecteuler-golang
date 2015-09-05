package combinatorics

import ()

/*
	In mathematics the nth central binomial coefficient is defined in terms of the binomial coefficient by

	2n Choose n = {(2n)!} / {(n!)^2}  for all n >= 0.
*/
func CentralBinomialCoefficient(n int64) int64 {
	var result int64 = 1
	for i := n + 1; i <= 2*n; i++ {
		result *= i
		result /= (i - n)
	}
	return result
}
