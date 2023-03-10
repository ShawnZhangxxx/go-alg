package main
//The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
//of rows like this: (you may want to display this pattern in a fixed font for
//better legibility)
//
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//
//
// And then read line by line: "PAHNAPLSIIGYIR"
//
// Write the code that will take a string and make this conversion given a
//number of rows:
//
//
//string convert(string s, int numRows);
//
//
//
// Example 1:
//
//
//Input: s = "PAYPALISHIRING", numRows = 3
//Output: "PAHNAPLSIIGYIR"
//
//
// Example 2:
//
//
//Input: s = "PAYPALISHIRING", numRows = 4
//Output: "PINALSIGYAHRPI"
//Explanation:
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
//
//
// Example 3:
//
//
//Input: s = "A", numRows = 1
//Output: "A"
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 1000
// s consists of English letters (lower-case and upper-case), ',' and '.'.
// 1 <= numRows <= 1000
//
// Related Topics 字符串 👍 1961 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
//找规律  先遍历行再遍历一个个周期 直接append进去
func convert(s string, numRows int) string {
	n,r := len(s),numRows
	if r==1 || r>=n {
		return  s
	}
	t := r*2 -2
	ans := make([]byte,0,n)
	for i:=0; i<r;i++ {//枚举矩阵的行
		for j:= 0;j+i < n ;j+=t {//枚举每个周期 第一个字符
			ans = append(ans,s[j+i])
			if 0<i && i< r - 1 && j+t < n { //第二个字符
				ans = append(ans,s[j+t-i])
			}
		}
	}

	return string(ans)
}
//leetcode submit region end(Prohibit modification and deletion)
