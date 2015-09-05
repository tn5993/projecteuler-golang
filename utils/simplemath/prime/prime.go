package prime

import ()

func SieveOfE(n int64) []bool {
	primeTable := make([]bool, n+1)
	for i := int64(2); i*i <= n; i++ {
		if !primeTable[i] {
			for j := i * i; j <= n; j += i {
				primeTable[j] = true
			}
		}
	}
	return primeTable
}

func GetPrimesUnder(n int64) []int64 {
	if n <= 1 {
		return []int64{}
	}

	primeTable := SieveOfE(n)

	var primes []int64
	for i, element := range primeTable {
		if i > 1 && !element {
			primes = append(primes, int64(i))
		}
	}

	return primes
}
