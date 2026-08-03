func countSeniors(details []string) int {
    seniors := 0
    for _, detail := range details {
        age := 0
        n1, _ := strconv.Atoi(string(detail[11]))
        n2, _ := strconv.Atoi(string(detail[12]))
        age += n1 * 10
        age += n2
        if age > 60 {
            seniors++
        }
    }

    return seniors
}

// 0-9 = phone
// 10 = gender
// 11-12 = age
// 13-14 seat