# hangul-go

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/aethiopicuschan/hangul-go.svg)](https://pkg.go.dev/github.com/aethiopicuschan/hangul-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/aethiopicuschan/hangul-go)](https://goreportcard.com/report/github.com/aethiopicuschan/hangul-go)
[![CI](https://github.com/aethiopicuschan/hangul-go/actions/workflows/ci.yaml/badge.svg)](https://github.com/aethiopicuschan/hangul-go/actions/workflows/ci.yaml)

This library provides operations for Hangul.

## Installation

```bash
go get -u github.com/aethiopicuschan/hangul-go
```

## Usage

```go
func main() {
	got, _ := hangul.Combine('ᄀ', 'ᅡ', 'ᆨ')
	fmt.Println(string(got))
	// Output: 각

	gotCho, gotJung, gotJong, _ := hangul.Split(got)
	fmt.Println(string(gotCho), string(gotJung), string(gotJong))
	// Output: ᄀ ᅡ ᆨ
}
```

### Notice

There may be cases where the rune obtained from `Split` is not displayed correctly due to the font or other related factors.
