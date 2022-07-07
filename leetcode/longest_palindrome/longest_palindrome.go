package longestpalindrome

// 中心扩散法
func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	begin := 0
	maxNum := 1

	for i := 0; i < n; i++ {
		// 由中间的一个字符开始扩散 babad
		left := i - 1
		right := i + 1
		for left >= 0 && right < n && s[left] == s[right] {
			if right-left+1 > maxNum {
				begin = left
				maxNum = right - left + 1
			}
			left--
			right++
		}
		// 由中间的2个字符开始扩散 baacd
		left = i
		right = i + 1
		for left >= 0 && right < n && s[left] == s[right] {
			if right-left+1 > maxNum {
				begin = left
				maxNum = right - left + 1
			}
			left--
			right++
		}
	}
	return s[begin : begin+maxNum]
}
