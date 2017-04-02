// Copyright © 2017 Nick Klauer <klauer@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package adr

import (
	"errors"
	"path"
)

// DefaultDir is the base directory that will be created for storing Architecture
// Decision Records.
const DefaultDir = "adr"
const pathNotFoundErr = "ADR Path Not Found"

var searchPaths = []string{"adr", "docs", "arch"}

// InitializeConfig creates the DefaultDir directory only if it's not already present
func InitializeConfig() (bool, error) {
	p, err := FindADRPath()
	if err != nil {
		err := appFs.MkdirAll(path.Join(p, DefaultDir), 0711)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return true, nil
}

// FindADRPath will look up the ADR path with the given defaults.
func FindADRPath() (string, error) {
	for _, v := range searchPaths {
		if f, err := appFs.Open(v); err == nil {
			if st, er := f.Stat(); er == nil {
				if st.IsDir() {
					return v, nil
				}
			}
			return "", errors.New(pathNotFoundErr)
		}
	}
	return "", errors.New(pathNotFoundErr)
}
