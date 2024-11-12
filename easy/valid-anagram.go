package easy

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    count := make(map[rune]int)

    for _, ch := range s {
        count[ch] += 1
    }

    for _, ch := range t {
        count[ch] -= 1

        if count[ch] < 0 {
            return false
        }
    }

    return true

}