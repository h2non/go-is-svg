package issvg

import (
	"testing"
)

func TestIsSVG(t *testing.T) {
	cases := []struct {
		buf     []byte
		matches bool
	}{
		{[]byte{0}, false},
		{[]byte{0x00, 0x01, 0x03, 0x04}, false},
		{[]byte("<svg></svg>"), true},
		{[]byte(`<svg width="100" height="100"><circle cx="50" cy="50" r="40" stroke="green" stroke-width="4" fill="yellow" /></svg>`), true},
		{[]byte(`<?xml version="1.0" encoding="utf-8"?>
						 <!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1 Basic//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11-basic.dtd">
						 <svg version="1.1" baseProfile="basic" id="svg2" xmlns:svg="http://www.w3.org/2000/svg" viewBox="0 0 900 900" xml:space="preserve">
						 	<path id="path482" fill="none" d="M184.013,144.428"/>
						 	<path id="path6" fill="#FFFFFF" stroke="#000000" stroke-width="0.172" d="917-66.752-80.957C40.928,326.18,72.326,313.197,108.956,403.826z"/>
						 </svg>`), true},
	}

	for _, test := range cases {
		if IsSVG(test.buf) != test.matches {
			t.Fatalf("Invalid SVG type: %s", test.buf)
		}
	}
}
