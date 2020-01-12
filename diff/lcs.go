package diff

import (
	"github.com/antzucaro/matchr"
)

// VanillaLCS computes the reserences between two strings using the LCS
// algortithm from https://en.m.wikipedia.org/wiki/Longest_common_subsequence_problem
func VanillaLCS(l, r []string) Result {
	return diffLCS(l, r)
}

func diffLCS(l, r []string, refiners ...Tokenizer) (res Result) {
	var wL, wR []string
	var resHead, resTail Result
	var m [][]int

	lookForSimilar := (len(refiners) > 0)

	wL, wR, resHead = getSameHead(l, r)
	wL, wR, resTail = getSameTail(wL, wR)

	if lookForSimilar {
		m = matrixLSS(wL, wR)
	} else {
		m = matrixLCS(wL, wR)
	}

	var iL, iR int
	for iL, iR = len(wL), len(wR); iL != 0 && iR != 0; {
		if m[iL][iR] == m[iL][iR-1] {
			res.insert(newInsertedDiff(wR[iR-1]))
			iR--
		} else if m[iL][iR] == m[iL-1][iR] {
			res.insert(newDeletedDiff(wL[iL-1]))
			iL--
		} else {
			if lookForSimilar {
				res.insert(adaptative(diffLCS, wL[iL-1], wR[iR-1], refiners...))
			} else {
				res.insert(newSameDiff(wL[iL-1]))
			}
			iL--
			iR--
		}
	}

	for iL != 0 {
		res.insert(newDeletedDiff(wL[iL-1]))
		iL--
	}

	for iR != 0 {
		res.insert(newInsertedDiff(wR[iR-1]))
		iR--
	}

	res.insert(resHead...)
	res.append(resTail...)
	return
}

func matrixLCS(a, b []string) [][]int {
	m := make([][]int, len(a)+1)
	for i := range m {
		m[i] = make([]int, len(b)+1)
	}

	for ia := range a {
		for ib := range b {
			if a[ia] == b[ib] {
				m[ia+1][ib+1] = m[ia][ib] + 1
			} else if m[ia+1][ib] > m[ia][ib+1] {
				m[ia+1][ib+1] = m[ia+1][ib]
			} else {
				m[ia+1][ib+1] = m[ia][ib+1]
			}
		}
	}

	return m
}

func matrixLSS(a, b []string) [][]int {
	m := make([][]int, len(a)+1)
	for i := range m {
		m[i] = make([]int, len(b)+1)
	}

	isSimilar := similar(a, b)

	for ia := range a {
		for ib := range b {
			if isSimilar[ia][ib] {
				m[ia+1][ib+1] = m[ia][ib] + 1
			} else if m[ia+1][ib] > m[ia][ib+1] { // MAX seq N-1
				m[ia+1][ib+1] = m[ia+1][ib]
			} else {
				m[ia+1][ib+1] = m[ia][ib+1]
			}
		}
	}

	return m
}

func sequenceLCS(a, b []string, lookForSimilar bool) (seq [][2]int) {
	var m [][]int

	if lookForSimilar {
		m = matrixLSS(a, b)
	} else {
		m = matrixLCS(a, b)
	}

	for ia, ib := len(a), len(b); ia != 0 && ib != 0; {
		if m[ia][ib] == m[ia-1][ib] {
			ia--
		} else if m[ia][ib] == m[ia][ib-1] {
			ib--
		} else {
			seq = append(seq, [2]int{ia - 1, ib - 1})
			ia--
			ib--
		}
	}

	for i, j := 0, len(seq)-1; i < j; i, j = i+1, j-1 {
		seq[i], seq[j] = seq[j], seq[i]
	}
	return seq
}

func getSameHead(l, r []string) ([]string, []string, Result) {
	var res Result

	var i int
	for i = 0; i < len(l) && i < len(r); i++ {
		if l[i] != r[i] {
			return l[i:], r[i:], res
		}
		res.append(newSameDiff(l[i]))
	}

	return l[i:], r[i:], res
}

func getSameTail(l, r []string) ([]string, []string, Result) {
	var res Result

	lenL, lenR := len(l), len(r)
	lastL, lastR := len(l)-1, len(r)-1

	var i int
	for i = 0; i < lenL && i < lenR; i++ {
		if l[lastL-i] != r[lastR-i] {
			return l[:lenL-i], r[:lenR-i], res
		}
		res.insert(newSameDiff(l[lastL-i]))
	}

	return l[:lenL-i], r[:lenR-i], res
}

func similar(a, b []string) [][]bool {
	dist := make([][]float64, len(a))
	for i := range dist {
		dist[i] = make([]float64, len(b))
	}

	isSame := make([][]bool, len(a))
	for i := range isSame {
		isSame[i] = make([]bool, len(b))
	}

	for ia := range a {

		var max float64
		for ib := range b {
			dist[ia][ib] = matchr.JaroWinkler(a[ia], b[ib], true)
			if dist[ia][ib] > max {
				max = dist[ia][ib]
			}
		}
		if max < 0.75 {
			continue
		}
		for ib := range b {
			isSame[ia][ib] = (dist[ia][ib] == max)
		}
	}

	return isSame
}
