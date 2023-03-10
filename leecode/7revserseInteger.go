package main

import "math"

//Given a signed 32-bit integer x, return x with its digits reversed. If
//reversing x causes the value to go outside the signed 32-bit integer range [-2³¹, 2³¹ -
// 1], then return 0.
//
// Assume the environment does not allow you to store 64-bit integers (signed
//or unsigned).
//
//
// Example 1:
//
//
//Input: x = 123
//Output: 321
//
//
// Example 2:
//
//
//Input: x = -123
//Output: -321
//
//
// Example 3:
//
//
//Input: x = 120
//Output: 21
//
//
//
// Constraints:
//
//
// -2³¹ <= x <= 2³¹ - 1
//
// Related Topics 数学 👍 3770 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x%10
		x /= 10
		rev = rev*10 + digit
	}
	return
}
//leetcode submit region end(Prohibit modification and deletion)
