package hangul

// Combine combines the constituent components of a Korean syllable.
func Combine(choseong, jungseong, jongseong rune) (syllable rune, err error) {
	LIndex := int(choseong - LBase)
	VIndex := int(jungseong - VBase)
	TIndex := 0

	if jongseong != 0 {
		TIndex = int(jongseong - TBase)
	}

	if LIndex < 0 || LIndex >= LCount || VIndex < 0 || VIndex >= VCount || TIndex < 0 || TIndex >= TCount {
		err = ErrInvalidJamo
		return
	}

	syllable = SBase + rune((LIndex*VCount+VIndex)*TCount+TIndex)
	return
}
