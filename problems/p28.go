package problems

/*
	Problem description: https://projecteuler.net/problem=28
*/

func NewProblem28(n int64) IProblem {
	return problem28{n}
}

type problem28 struct {
	n int64
}

func (p problem28) Solve() int64 {
	if p.n%2 == 0 {
		return -1
	}

	var sum int64 = 1
	for i := int64(3); i <= 1001; i += 2 {
		sum += (4 * i * i) - 6*(i-1)
	}

	return sum
}

/*
	Solution:
	Let define a circle as the four corners in a n*n spiral
	Let n be the side length where 0 < n <= 1001. We have the four corner numbers as below
	n=1: 1
	n=3: 3 5 7 9
	n=5: 13 17 21 25
	n=7: 31 37 43 49
	...
	Notice that n is odd
	Except n=1, we see that the last number of a circle is n^2 because it is the area of a square. Then we see that
	the third number is n^2 - (n-1); the second number is n^2-2(n-1); the first number is n^2-3(n-1). Therefore, the
	sum of all numbers in a circle can be written as f(n) = 4n^2-6(n-1).

	f(1) = 1
	f(3) = 24
	f(5) = 76
	f(7) = 160
	f(9) = 324-48= 276
	f(11) = 484-60= 424
	f(13) = 676-72= 604

	To get the sum of all the number, we can loop use loop through from n=1 to n=1001. However, as eager to find a closed
	form equation for this problem (or because I hate for loop), I would dig deeper to find a general equation for the sum




	Using differences table, we have
		 	 n			3				5				7 			9			11		13
		f(n) 	 	 24			 76			160			276		 424	 604
delta(n1)					   52			 84			116		 148   180
delta(n2)						 				 32			 32		  32		32

	Because the constant value appears at the second difference, we conclude that the equation producing our values is of degree 2 (quadratic equation)
	f(x) = a*x^2 + b*x + c

	According to Ken Ward's mathematics page, we can obtains the following equation for the quadratic function
*/
