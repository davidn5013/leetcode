package grind75

/* Climb stairs problem: Two version recursive and iterative

Bench:
goos: windows
goarch: amd64
pkg: toBeSorted/test/2973525193
cpu: AMD Ryzen 7 3700X 8-Core Processor
BenchmarkClimbStairsRecursive
BenchmarkClimbStairsRecursive-16             253           4690272 ns/op
BenchmarkClimbStairs
BenchmarkClimbStairs-16                      505           2370501 ns/op
PASS
ok      toBeSorted/test/2973525193      3.264s
*/

// ClimbStairs 0070 in how many distinct ways can you climb to the top of n stairs? This a iterativel fast version.
/*
The number of ways to climb the stairs is equal to the sum of the number of ways to climb the
stairs from the previous step and the number of ways to climb the stairs from two steps ago.

In other words, if we let f(n) be the number of ways to climb the stairs
starting from the nth step, then we have the recurrence:
f(n) = f(n - 1) + f(n - 2)

We can start the recursion with the base cases:
f(1) = 1 (there is only 1 way to climb 1 step)
f(2) = 2 (there are 2 ways to climb 2 steps: 1+1 or 2)

Then, we can use the recurrence to compute the number of ways to climb the stairs for n steps:

f(n) = f(n - 1) + f(n - 2)
= f(n - 2) + f(n - 3) + f(n - 2)
= f(n - 3) + f(n - 4) + f(n - 2) + f(n - 2)
= ...

This is known as the Fibonacci sequence, and it has the closed-form solution:
f(n) = ((1 + sqrt(5))/2)^n - ((1 - sqrt(5))/2)^n
where sqrt(5) is the square root of 5.
*/
func ClimbStairs(n int) int {
	switch {
	case n < 1:
		return 0
	case n == 1 || n == 2:
		return n
	}
	f1, f2 := 1, 2
	for i := 3; i <= n; i++ {
		f := f1 + f2
		f1, f2 = f2, f
	}
	return f2
}

// ClimbStairsRecursive This the slow recursive version! In how many distinct ways can you climb to the top of n stairs?
func ClimbStairsRecursive(n int) int {
	switch {
	case n < 1:
		return 0
	case n == 1 || n == 2:
		return n
	}
	//
	return ClimbStairs(n-1) + ClimbStairs(n-2)
}
