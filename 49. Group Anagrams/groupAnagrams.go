package _9__Group_Anagrams

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	dic := map[string][]string{}
	for _, s := range strs {
		sortedString := SortString(s)
		_, ok := dic[sortedString]
		if ok {
			dic[sortedString] = append(dic[sortedString], s)
		} else {
			dic[sortedString] = []string{s}
		}
	}
	res := [][]string{}
	for _, i := range dic {
		res = append(res, i)
	}

	return res

}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
