package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/**
 * Author       :   %s
 * Date         :   %s
 * ver.         :   %s
 * link         :   %s
 * dir          :   %s
 */

const MOD = 1_000_000_007 // 10^9+7

var (
	reader *bufio.Reader
	writer *bufio.Writer
)

func init() {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}

func code(a []int, n int) int {
	return 0
}

func main() {
	defer writer.Flush()

	// const TOTAL_TEST_CASES = 1
	TOTAL_TEST_CASES := nextInt()

	for testCase := 0; testCase < TOTAL_TEST_CASES; testCase++ {
		// Taking input
		// n := nextInt()
		// s := next()
		// a := inputArray(n)

		// Compute
		// res := code(a, n)
		// fmt.Fprintf(writer, "DEBUG: Case #%d: %d\n", testCase+1, res)

		// Storing the result
		// fmt.Fprintln(writer, res)
	}
}

// ==================== Input Helper Functions ====================

func next() string {
	var token string
	fmt.Fscan(reader, &token)
	return token
}

func nextInt() int {
	var n int
	fmt.Fscan(reader, &n)
	return n
}

func nextInt64() int64 {
	var n int64
	fmt.Fscan(reader, &n)
	return n
}

func nextFloat64() float64 {
	var f float64
	fmt.Fscan(reader, &f)
	return f
}

func nextLine() string {
	line, _ := reader.ReadString('\n')
    return strings.TrimRight(line, "\r\n")
}

func inputArray(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	return a
}

func inputArray64(n int) []int64 {
	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt64()
	}
	return a
}

func inputMatrix(r, c int) [][]int {
	matrix := make([][]int, r)
	for i := 0; i < r; i++ {
		matrix[i] = inputArray(c)
	}
	return matrix
}

// ==================== Utility Functions ====================

func displayMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}

func modPow(base, exp, mod int) int {
	result := 1
	base %= mod
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}
	return result
}

func sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}

func reverseArray(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// ==================== String Conversion ====================

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func toInt64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

func toString(n int) string {
	return strconv.Itoa(n)
}

// ==================== Math Helpers ====================

func sqrtInt(n int) int {
	return int(math.Sqrt(float64(n)))
}

func ceilDiv(a, b int) int {
	return (a + b - 1) / b
}
