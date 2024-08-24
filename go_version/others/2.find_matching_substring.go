// 主串S，长度m(由26个字母组成)。模式串K，长度n。
// 匹配过程中根据子串的权值去匹配，例如“ABC”的权值为100+20+3，去匹配主串。
// 权值相同则匹配。写出相应的程序。

package others

import "math"

func calculateWeight(s string) int {
	weight := 0
	n := len(s)
	for i, c := range s {
		weight += int(c-'a'+1) * int(math.Pow(10, float64(n-i-1)))
	}
	return weight
}

func findMatchingSubstrings(s, k string) []int {
	n := len(k)
	matches := make([]int, 0)

	// 计算模式串 K 的权值
	weightK := calculateWeight(k)

	// 遍历主串 S 的所有长度为 n 的子串
	for i := 0; i <= len(s)-n; i++ {
		subStr := s[i : i+n]
		weightS := calculateWeight(subStr)

		if weightS == weightK {
			matches = append(matches, i)
		}
	}
	return matches
}
