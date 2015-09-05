package intmath

import (
	"errors"
	"math/big"
)

func Max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func Max64(a, b int64) int64 {
	if a < b {
		return b
	}

	return a
}

func Abs(a int) int {
	if a < 0 {
		a = (-1) * a
	}
	return a
}

func Factorial(n int64) (int64, error) {
	if n < 0 {
		return -1, errors.New("Factorial does not apply for negative number")
	}

	var fac int64 = 1
	for i := int64(2); i <= n; i++ {
		fac *= i
	}

	return fac, nil
}

func BigFactorial(n int) (*big.Int, error) {
	if n < 0 {
		return big.NewInt(-1), errors.New("Factorial does not apply for negative number")
	}

	fac := big.NewInt(1)
	for i := 2; i <= n; i++ {
		fac.Mul(fac, big.NewInt(int64(i)))
	}

	return fac, nil
}

func SumDigit(num *big.Int) *big.Int {
	zeroInt := big.NewInt(0)
	tenInt := big.NewInt(10)
	sumInt := big.NewInt(0)
	modInt := new(big.Int)
	for num.Cmp(zeroInt) > 0 {
		modInt.Mod(num, tenInt)
		sumInt.Add(sumInt, modInt)
		num.Div(num, tenInt)
	}

	return sumInt
}
