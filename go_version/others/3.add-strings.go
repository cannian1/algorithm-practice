// 415.字符串相加
// https://leetcode.cn/problems/add-strings/

package others

import "strconv"

func addStrings(num1 string, num2 string) string {
	res := ""
	carry := 0 // 进位
	i, j := len(num1)-1, len(num2)-1

	for i >= 0 || j >= 0 || carry > 0 {
		var digit1, digit2 int
		if i >= 0 {
			digit1 = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			digit2 = int(num2[j] - '0')
			j--
		}
		sum := digit1 + digit2 + carry
		carry = sum / 10
		res = strconv.Itoa(sum%10) + res
	}
	return res
}
