package diff

type algorithm func([]string, []string, ...Tokenizer) Result

func adaptative(differ algorithm, l, r string, tokenizers ...Tokenizer) Result {
	if len(tokenizers) == 0 {
		return nil
	}

	wL, wR := tokenizers[0](l), tokenizers[0](r)
	res := differ(wL, wR)

	if len(tokenizers) == 1 {
		return res
	}

	// for zones that are differents (serie of deletion and insertions), we
	// look for similar areas for which we try to refine the differences.
	zones := res.differentZones()
	if len(zones) == 0 {
		return res
	}

	rdiff := res[:zones[0][0]]
	for i := range zones {
		aL, aR := res[zones[i][0] : zones[i][1]+1].content()
		adiff := differ(aL, aR, tokenizers[1:]...)
		if len(adiff) != 0 {
			rdiff = append(rdiff, adiff...)
		}

		if i == len(zones)-1 {
			if zones[i][1] < len(res)-1 {
				rdiff = append(rdiff, res[zones[i][1]+1:]...)
			}
		} else {
			rdiff = append(rdiff, res[zones[i][1]+1:zones[i+1][0]]...)
		}
	}

	return rdiff
}
