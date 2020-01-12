package diff

// LCS computes the differences between l and r strings using the LCS
// algorithm.
func LCS(l, r string, tokenizers ...Tokenizer) Result {
	if len(tokenizers) == 0 {
		tokenizers = []Tokenizer{ByLines}
	}

	return adaptative(diffLCS, l, r, tokenizers...)
}

// Patience computes the differences between l and r strings using the Patience
// algorithm.
func Patience(l, r string, tokenizers ...Tokenizer) Result {
	if len(tokenizers) == 0 {
		tokenizers = []Tokenizer{ByLines}
	}

	return adaptative(diffPatience, l, r, tokenizers...)
}
