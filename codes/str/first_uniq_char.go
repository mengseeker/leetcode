// 字符串中的第一个唯一字符
// 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

//

// 示例：

// s = "leetcode"
// 返回 0

// s = "loveleetcode"
// 返回 2
//

// 提示：你可以假定该字符串只包含小写字母。

// 作者：力扣 (LeetCode)
// 链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn5z8r/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

package str

func firstUniqChar(s string) int {
	var chars [128]int
	for _, c := range s {
		chars[c]++
	}
	for i, c := range s {
		if chars[c] == 1 {
			return i
		}
	}
	return -1
}
