package medium

import "sort"

func groupAnagrams(strs []string) [][]string {
    annMap := make(map[string][]string)
    signature := ""
    for _, s := range strs {
        signature = getSignature(s)

        annMap[signature] = append(annMap[signature], s)
    }

    result := make([][]string, 0, len(annMap))
    for _, grp := range annMap {
        result = append(result, grp)
    }

    return result
}

func getSignature(s string) string{
    ch := []rune(s)
    sort.Slice(ch, func(i int, j int) bool {
        return ch[i] < ch[j]
    })
    return string(ch)
}
