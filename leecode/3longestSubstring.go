package main

import "fmt"

//Given a string s, find the length of the longest substring without repeating
//characters.
//
//
// Example 1:
//
//
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//
//
// Example 2:
//
//
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//
//
// Example 3:
//
//
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a
//substring.
//
//
//
// Constraints:
//
//
// 0 <= s.length <= 5 * 10⁴
// s consists of English letters, digits, symbols and spaces.
//
// Related Topics 哈希表 字符串 滑动窗口 👍 8725 👎 0
//func main()  {
//	lengthOfLongestSubstring("abdabdaa")
//}
//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	//
	m := map[byte]int{}
	n := len(s)
	fmt.Println(m[69])
	//右指针,初始值为-1 ,
	rk,ans := -1 ,0
	fmt.Println(ans)

	for i:= 0 ; i<n ;i++ {
		if i != 0 {
			//左指针向右移动一格,移除一个字符
			delete(m,s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {  //map 未赋值的初始值为0
			//不断移动指针
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans,rk -i + 1)
	}
	return  ans
}

func max(x,y int) int {
	if x > y {
		return x
	}else {
		return y
	}
}

//leetcode submit region end(Prohibit modification and deletion)
