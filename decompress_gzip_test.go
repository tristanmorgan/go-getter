// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package getter

import (
	"path/filepath"
	"testing"
)

func TestGzipDecompressor(t *testing.T) {
	cases := []TestDecompressCase{
		{
			"single.gz",
			false,
			false,
			nil,
			"d3b07384d113edec49eaa6238ad5ff00",
			nil,
		},

		{
			"single.gz",
			true,
			true,
			nil,
			"",
			nil,
		},
	}

	for i, tc := range cases {
		cases[i].Input = filepath.Join("./testdata", "decompress-gz", tc.Input)
	}

	TestDecompressor(t, new(GzipDecompressor), cases)
}
