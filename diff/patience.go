package diff

// VanillaPatience implements the patience diff algorithm (from
// http://alfedenzo.livejournal.com/170301.html)
func VanillaPatience(l, r []string) (res Result) {
	return diffPatience(l, r)
}

func diffPatience(l, r []string, refiners ...Tokenizer) (res Result) {
	var wL, wR []string
	var resHead, resTail Result

	lookForSimilar := (len(refiners) > 0)

	wL, wR, resHead = getSameHead(l, r)
	wL, wR, resTail = getSameTail(wL, wR)

	uniqL, uniqR := newUniq(wL), newUniq(wR)

	var iL, iR int
	for _, seq := range sequenceLCS(uniqL.Content, uniqR.Content, lookForSimilar) {
		iendL, iendR := uniqL.OrigPos[seq[0]], uniqR.OrigPos[seq[1]]
		res.append(diffLCS(wL[iL:iendL], wR[iR:iendR], refiners...)...)
		if lookForSimilar {
			res.append(adaptative(diffLCS, wL[iendL], wR[iendR], refiners...))
		} else {
			res.append(newSameDiff(wL[iendL]))
		}
		iL, iR = iendL+1, iendR+1
	}

	res.append(diffLCS(wL[iL:], wR[iR:], refiners...)...)
	res.insert(resHead...)
	res.append(resTail...)
	return
}

type uniq struct {
	Content []string
	OrigPos []int
}

func newUniq(a []string) *uniq {
	u := &uniq{}
	for i, r := range a {
		if pos := u.Pos(r); pos >= 0 {
			u.Del(pos)
		} else {
			u.Add(r, i)
		}
	}
	return u
}

func (u *uniq) Add(r string, pos int) {
	u.Content = append(u.Content, r)
	u.OrigPos = append(u.OrigPos, pos)
}

func (u *uniq) Del(pos int) {
	u.Content = append(u.Content[:pos], u.Content[pos+1:]...)
	u.OrigPos = append(u.OrigPos[:pos], u.OrigPos[pos+1:]...)
}

func (u *uniq) Pos(r string) int {
	for i, c := range u.Content {
		if c == r {
			return i
		}
	}
	return -1
}
