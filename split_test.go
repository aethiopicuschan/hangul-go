package hangul_test

import (
	"fmt"
	"testing"

	"github.com/aethiopicuschan/hangul-go"
	"github.com/stretchr/testify/assert"
)

func ExampleSplit() {
	gotCho, gotJung, gotJong, _ := hangul.Split('각')
	fmt.Println(string(gotCho), string(gotJung), string(gotJong))
	// Output: ᄀ ᅡ ᆨ
}

func TestSplit(t *testing.T) {
	tests := []struct {
		name      string
		syllable  rune
		choseong  rune
		jungseong rune
		jongseong rune
		expectErr bool
	}{
		{
			name:      "ga (ᄀ + ᅡ)",
			syllable:  '가', // U+AC00
			choseong:  'ᄀ', // U+1100
			jungseong: 'ᅡ', // U+1161
			jongseong: 0,
		},
		{
			name:      "gak (ᄀ + ᅡ + ᆨ)",
			syllable:  '각', // U+AC01
			choseong:  'ᄀ',
			jungseong: 'ᅡ',
			jongseong: 'ᆨ', // U+11A8
		},
		{
			name:      "han (ᄒ + ᅡ + ᆫ)",
			syllable:  '한', // U+D55C
			choseong:  'ᄒ',
			jungseong: 'ᅡ',
			jongseong: 'ᆫ',
		},
		{
			name:      "o (ᄋ + ᅩ)",
			syllable:  '오', // U+C624
			choseong:  'ᄋ', // U+110B
			jungseong: 'ᅩ', // U+1169
			jongseong: 0,
		},
		{
			name:      "Error: Invalid Syllable (A)",
			syllable:  'A',
			expectErr: true,
		},
		{
			name:      "Error: Invalid Syllable (non-Hangul)",
			syllable:  '🙂', // Emoji (non-Hangul character)
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			choseong, jungseong, jongseong, err := hangul.Split(tt.syllable)

			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.choseong, choseong)
			assert.Equal(t, tt.jungseong, jungseong)
			assert.Equal(t, tt.jongseong, jongseong)
		})
	}
}
