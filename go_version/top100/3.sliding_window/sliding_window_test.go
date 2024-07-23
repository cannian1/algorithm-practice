package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	str1 := "abcabcbb"
	str2 := "bbbbb"
	str3 := "pwwkew"

	r1 := lengthOfLongestSubstring(str1)
	r2 := lengthOfLongestSubstring(str2)
	r3 := lengthOfLongestSubstring(str3)

	assert.Equal(t, r1, 3)
	assert.Equal(t, r2, 1)
	assert.Equal(t, r3, 3)
}
