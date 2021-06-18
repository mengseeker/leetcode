// 验证回文串
// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

// 说明：本题中，我们将空字符串定义为有效的回文串。

// 示例 1:

// 输入: "A man, a plan, a canal: Panama"
// 输出: true
// 示例 2:

// 输入: "race a car"
// 输出: false
// 相关标签

// 作者：力扣 (LeetCode)
// 链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xne8id/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

package str

import (
	"bytes"
)

func isPalindrome(s string) bool {
	bs := bytes.ToUpper([]byte(s))
	var i, j = 0, len(bs) - 1
	for !((bs[i] >= '0' && bs[i] <= '9') || (bs[i] >= 'a' && bs[i] <= 'z') || (bs[i] >= 'A' && bs[i] <= 'Z')) {
		i++
		if i >= len(bs) {
			return true
		}
	}
	for !((bs[j] >= '0' && bs[j] <= '9') || (bs[j] >= 'a' && bs[j] <= 'z') || (bs[j] >= 'A' && bs[j] <= 'Z')) {
		j--
	}
	for i < j {
		if bs[i] != bs[j] {
			return false
		}
		i++
		j--
		for !((bs[i] >= '0' && bs[i] <= '9') || (bs[i] >= 'a' && bs[i] <= 'z') || (bs[i] >= 'A' && bs[i] <= 'Z')) {
			i++
		}
		for !((bs[j] >= '0' && bs[j] <= '9') || (bs[j] >= 'a' && bs[j] <= 'z') || (bs[j] >= 'A' && bs[j] <= 'Z')) {
			j--
		}
	}
	return true
}
