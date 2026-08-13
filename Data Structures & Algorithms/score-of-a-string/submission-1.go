func scoreOfString(s string) int {
    score := 0
    for i, _ := range s {
        if i+1 == len(s) {
            continue
        }
        temp := int(s[i]) - int(s[i+1])
        if temp < 0 {
            temp *= -1
        }

        score += temp
    }
    return score
}

