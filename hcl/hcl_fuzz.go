// +build gofuzz

package hcl

import hhcl "github.com/hashicorp/hcl"

// A return value of 0 from Fuzz() means the data wasn’t interesting — the parser detected an error.
// A return value of 1 means the data was parsed successfully, even though it had been modified.
// The fuzzer keeps these as more interesting inputs, since they appear to be valid.

func Fuzz(data []byte) int {
	if _, err := hhcl.ParseBytes(data); err != nil {
		return 0
	}
	return 1
}
