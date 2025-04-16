package hangul

// Split splits a Korean syllable into its constituent components.
func Split(syllable rune) (choseong, jungseong, jongseong rune, err error) {
	if syllable < SBase || syllable > SBase+(LCount*VCount*TCount)-1 {
		err = ErrInvalidSyllable
		return
	}

	syllableIndex := syllable - SBase

	LIndex := syllableIndex / (VCount * TCount)
	VIndex := (syllableIndex % (VCount * TCount)) / TCount
	TIndex := syllableIndex % TCount

	choseong = LBase + rune(LIndex)
	jungseong = VBase + rune(VIndex)
	if TIndex != 0 {
		jongseong = TBase + rune(TIndex)
	} else {
		jongseong = 0
	}

	return
}
