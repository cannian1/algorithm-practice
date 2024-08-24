// 1. go 语言实现 10进制转 32进制

package others

const base32Chars = "0123456789abcdefghijklmnopqrstuv" // 32个字符表示0-31

func DecToBase32(num int) string {
	if num == 0 {
		return "0"
	}

	base := 32
	result := ""

	for num > 0 {
		remainder := num % base
		result = string(base32Chars[remainder]) + result
		num /= base
	}

	return result
}
