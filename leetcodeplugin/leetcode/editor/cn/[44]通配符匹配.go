//给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。 
//
// '?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符串（包括空字符串）。
// 
//
// 两个字符串完全匹配才算匹配成功。 
//
// 说明: 
//
// 
// s 可能为空，且只包含从 a-z 的小写字母。 
// p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。 
// 
//
// 示例 1: 
//
// 输入:
//s = "aa"
//p = "a"
//输出: false
//解释: "a" 无法匹配 "aa" 整个字符串。 
//
// 示例 2: 
//
// 输入:
//s = "aa"
//p = "*"
//输出: true
//解释: '*' 可以匹配任意字符串。
// 
//
// 示例 3: 
//
// 输入:
//s = "cb"
//p = "?a"
//输出: false
//解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。
// 
//
// 示例 4: 
//
// 输入:
//s = "adceb"
//p = "*a*b"
//输出: true
//解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".
// 
//
// 示例 5: 
//
// 输入:
//s = "acdcb"
//p = "a*c?b"
//输出: false 
// Related Topics 贪心算法 字符串 动态规划 回溯算法 
// 👍 559 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isMatch(s string, p string) bool {
if (len(p) == 0 || p =="*" ) && len(s) == 0 {
return true
}

if len(p) == 0 || (len(s) == 0 && p[0]!='*') {
return false
}

switch p[0] {
case '*':
tag := 0
stag := 0
for i := 1; i < len(p); i++ {
if p[i] == '*' {
continue
}
if p[i] == '?' {
stag++
continue
}
tag = i
break
}


res :=  isMatch("", p[tag:])
for i := stag; i < len(s); i++ {
if s[i] == p[tag] {
res = res || isMatch(s[i:], p[tag:])
}
}
return res
case '?':
default:
if p[0] != s[0] {
return false
}
}
return isMatch(s[1:], p[1:])
}
//leetcode submit region end(Prohibit modification and deletion)
