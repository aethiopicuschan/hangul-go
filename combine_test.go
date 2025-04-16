package hangul_test

import (
	"fmt"
	"testing"

	"github.com/aethiopicuschan/hangul-go"
	"github.com/stretchr/testify/assert"
)

func ExampleCombine() {
	got, _ := hangul.Combine('ᄀ', 'ᅡ', 'ᆨ')
	fmt.Println(string(got))
	// Output: 각
}

func TestCombineHangul(t *testing.T) {
	tests := []struct {
		name      string
		choseong  rune
		jungseong rune
		jongseong rune
		want      rune
		expectErr bool
	}{
		{
			name:      "ga (ᄀ + ᅡ)",
			choseong:  'ᄀ', // U+1100
			jungseong: 'ᅡ', // U+1161
			jongseong: 0,
			want:      '가', // U+AC00
		},
		{
			name:      "gak (ᄀ + ᅡ + ᆨ)",
			choseong:  'ᄀ',
			jungseong: 'ᅡ',
			jongseong: 'ᆨ', // U+11A8
			want:      '각', // U+AC01
		},
		{
			name:      "han (ᄒ + ᅡ + ᆫ)",
			choseong:  'ᄒ',
			jungseong: 'ᅡ',
			jongseong: 'ᆫ',
			want:      '한',
		},
		{
			name:      "o (ᄋ + ᅩ)",
			choseong:  'ᄋ',
			jungseong: 'ᅩ',
			jongseong: 0,
			want:      '오',
		},
		{
			name:      "Error: Invalid Choseong",
			choseong:  'A',
			jungseong: 'ᅡ',
			jongseong: 0,
			expectErr: true,
		},
		{
			name:      "Error: Invalid Jungseong",
			choseong:  'ᄀ',
			jungseong: 'A',
			jongseong: 0,
			expectErr: true,
		},
		{
			name:      "Error: Invalid Jongseong",
			choseong:  'ᄀ',
			jungseong: 'ᅡ',
			jongseong: 'A',
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := hangul.Combine(tt.choseong, tt.jungseong, tt.jongseong)

			if tt.expectErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
