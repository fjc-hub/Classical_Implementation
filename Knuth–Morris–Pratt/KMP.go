/**
 * 	1. Knuth-Morris-Pratt substring search, A Clever method to always avoid text pointer backup!!!
 * 	2. Deterministic finite state automaton(DFA, not NFA) implementation.
 *  3. It finds the first index of the substring in Text string,
 *     such that this substring is same as Pattern, if not exist, return -1.
 *  4. Alphabet: a-z, support lower case only
 *	5. Testing: https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/submissions/379102145/
 *  6. Reference: https://d3c33hcgiwev3.cloudfront.net/_551fa6c1c84ff226fc5ccf79ec29cf0b_53SubstringSearch.pdf?Expires=1667606400&Signature=bCdUZ3r64zDNDwagTNBj8m6JXK4eSiSkTEK2-pSL5IqsTU4lqRe0yIfk8Ul8-25M3RSzQKYkDtl-0~n9BPVasLHZBqkx~DciRwSMHVYn5xrj2uysnmSmOnVtabZP0MBPzHI-KioyFNkVpmNTULOyTNqVEe5bNYUjKtH87T847FI_&Key-Pair-Id=APKAJLTNE6QMUY6HBC5A
 */

package KMP

func KMP(text, pattern string) int {
	// get DFA
	dfa := constructDFA(pattern)
	// "simulate" on DFA
	state, i := 0, 0
	for ; i < len(text) && state < len(pattern); i++ { // text pointer i no backup
		state = dfa[state][text[i]-'a']
	}
	if state == len(pattern) {
		return i - len(pattern)
	}
	return -1 // not found
}

// construct DFA
func constructDFA(pattern string) [][]int {
	n := len(pattern)
	// init
	dfa := make([][]int, n)
	for i := range dfa {
		dfa[i] = make([]int, 26)
	}

	// state X is the result state of simulating pattern[1,i-1] on DFA
	X := 0

	// dynamic programming
	dfa[0][int(pattern[0]-'a')] = 1
	// count i from 1, because pattern[1,-1] is invalid
	for i := 1; i < n; i++ {
		t := int(pattern[i] - 'a')
		for j := 0; j < 26; j++ {
			if t == j {
				dfa[i][j] = i + 1
			} else {
				dfa[i][j] = dfa[X][j]
			}
		}
		X = dfa[X][int(pattern[i]-'a')]
	}
	return dfa
}
