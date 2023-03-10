package main
//Given an integer x, return true if x is a palindrome, and false otherwise.
//
//
// Example 1:
//
//
//Input: x = 121
//Output: true
//Explanation: 121 reads as 121 from left to right and from right to left.
//
//
// Example 2:
//
//
//Input: x = -121
//Output: false
//Explanation: From left to right, it reads -121. From right to left, it
//becomes 121-. Therefore it is not a palindrome.
//
//
// Example 3:
//
//
//Input: x = 10
//Output: false
//Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
//
//
//
// Constraints:
//
//
// -2³¹ <= x <= 2³¹ - 1
//
//
//
//Follow up: Could you solve it without converting the integer to a string?
//Related Topics 数学 👍 2430 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return  false
	}

	revertedNumber := 0
	for x > revertedNumber { //反转整数 和第7题差不多
		revertedNumber = revertedNumber * 10 + x % 10
		x /= 10
	}

	return  x == revertedNumber || x== revertedNumber / 10
}


//leetcode submit region end(Prohibit modification and deletion)

