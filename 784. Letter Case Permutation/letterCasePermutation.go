package _84__Letter_Case_Permutation

import "unicode"

func letterCasePermutation(S string) []string {
	var res []string
	for i, c := range S {
		if unicode.IsLetter(c) {
			for _, r := range letterCasePermutation(S[i+1:]) {
				res = append(res, S[:i]+string(unicode.ToUpper(c))+r)
				res = append(res, S[:i]+string(unicode.ToLower(c))+r)
			}
			break
		}
	}
	if len(res) == 0 {
		return []string{S}
	}

	return res
}
