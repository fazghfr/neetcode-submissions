func longestPalindrome(s string) string {
    maxLen := 0
    var l, r int
    for i, _ := range s {
        evLen := lengthPalindromeFromMiddle(s, i, i+1)
        oddLen := lengthPalindromeFromMiddle(s, i, i)

        if oddLen[0] > maxLen {
            maxLen = oddLen[0]
            l = oddLen[1]
            r = oddLen[2]
        }

        if evLen[0] > maxLen {
            maxLen = evLen[0]
            l = evLen[1]
            r = evLen[2]
        }
    }

    return s[l : r+1]
}

func lengthPalindromeFromMiddle(s string, left int, right int) []int {
    for left >= 0 && right < len(s) && s[left] == s[right] { 
        left--
        right++
    }

    left++ 
    right-- 

    ans := make([]int, 3)
    ans[0] = right - left + 1
    ans[1] = left
    ans[2] = right
    return ans
}