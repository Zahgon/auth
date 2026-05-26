// Package envparse is a fork of the github.com/joho/godotenv parser.
//
// This fork is based on master which has some minor fixes[1] since the v1.5.1
// we previously used.
//
// [1] https://github.com/joho/godotenv/compare/v1.5.1...main
//
// -------
//
// # Copyright (c) 2013 John Barton
//
// # MIT License
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
package envparse

import (
	"io"
)

// Parse reads an env file from io.Reader, returning a map of keys and values.
func Parse(r io.Reader) (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }

// ParseData is like Parse but works on a string or []byte slice.
func ParseData[T ~string | ~[]byte](data T) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// since we use data as scratch space

// parseData will mutate data during parsing, use ParseData to avoid this.
func parseData(data []byte) (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }

const (
	charComment       = '#'
	prefixSingleQuote = '\''
	prefixDoubleQuote = '"'

	exportPrefix = "export"
)

func parseBytes(src []byte, out map[string]string) error { _ = "STUB: not implemented"; return nil }

// reached end of file

// getStatementPosition returns position of statement begin.
//
// It skips any comment line or non-whitespace character.
func getStatementStart(src []byte) []byte { _ = "STUB: not implemented"; return nil }

// skip comment section

// locateKeyName locates and parses key name and returns rest of slice
func locateKeyName(src []byte) (key string, cutset []byte, err error) {
	_ = "STUB: not implemented"
	// trim "export" and space at beginning
	return "", nil, nil
}

// locate key name end and validate it in single loop

// library also supports yaml-style value declaration

// variable name should match [A-Za-z0-9_.]

// trim whitespace

// expandDollarEscapes preserves godotenvs dollar escaping.
func expandDollarEscapes(src []byte) []byte { _ = "STUB: not implemented"; return nil }

// extractVarValue extracts variable value and returns rest of slice.
func extractVarValue(src []byte) (value string, rest []byte, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// unquoted value - read until end of line

// Hit EOF without a trailing newline

// Convert line to rune away to do accurate countback of runes

// Assume end of line is end of var

// Strip trailing comments only when '#' is preceded by whitespace:
// FOO=bar # comment => "bar"
// FOO=bar#baz       => "bar#baz"
// FOO=#bar          => "#bar"

// lookup quoted string terminator

func isEscaped(src []byte, index int) bool { _ = "STUB: not implemented"; return false }

func expandEscapes(src []byte) []byte { _ = "STUB: not implemented"; return nil }

// TODO(cstockton): We keep '$' here for stricter compat with todays
// config. If we want to be more strict (e.g. \$ -> \$) we can emit
// the additional \\ as well.

// Preserve upstream godotenv behavior for non-dollar escapes:
// \" => ", \\ => \, \x => x.

func indexOfNonSpaceChar(src []byte) int { _ = "STUB: not implemented"; return 0 }

// hasQuotePrefix reports whether charset starts with single or double quote and returns quote character
func hasQuotePrefix(src []byte) (prefix byte, isQuoted bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func isCharFunc(char rune) func(rune) bool { _ = "STUB: not implemented"; return nil }

// isSpace reports whether the rune is a space character but not line break character
//
// this differs from unicode.IsSpace, which also applies line break as space
func isSpace(r rune) bool { _ = "STUB: not implemented"; return false }

func isLineEnd(r rune) bool { _ = "STUB: not implemented"; return false }
