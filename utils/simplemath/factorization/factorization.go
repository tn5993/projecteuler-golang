package factorization

import ()

func DivisorsCount(n int64) int64 {
	var result, k int64 = 1, n
	for i := int64(2); i <= k; i++ {
		var p int64 = 0
		for k%i == 0 {
			p++
			k /= i
		}
		result *= (p + 1)
	}
	return result
}

func DivisorSum(n int64) int64 {
	if n == 1 {
		return 1
	}
	var sum int64 = 1
	var k int64 = n

	for i := int64(2); i <= k; i++ {
		var p int64 = 1
		for k%i == 0 {
			p *= i
			k /= i
		}
		sum *= (p*i - 1) / (i - 1)
	}
	return sum - n
}
