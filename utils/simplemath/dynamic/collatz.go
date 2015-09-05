/*

Longest Collatz sequence
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/
package dynamic

import ()

func FindLongestCollatzSequence(limit int64) int64 {
	mem := make([]int64, limit+1)
	mem[1] = 1
	max := mem[1]
	numberWithLongestChain := int64(1)
	for i := int64(2); i <= limit; i++ {
		n := i
		k := int64(1)

		for n != 1 && n >= i {
			k += 1
			if n%2 == 0 {
				n = n / 2
			} else {
				n = 3*n + 1
			}
		}

		mem[i] = k + mem[n]
		if max < mem[i] {
			max = mem[i]
			numberWithLongestChain = i
		}
	}

	return numberWithLongestChain
}
