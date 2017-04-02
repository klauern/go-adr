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
	"bufio"
	"bytes"
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/spf13/afero"
)

// DecisionRecord is a type of Architecture Decision Record.
type DecisionRecord struct {
	Number       int       `json:"number,omitempty"`
	Title        string    `json:"title,omitempty"`
	Date         time.Time `json:"date,omitempty"`
	Status       string    `json:"status,omitempty"`
	Context      string    `json:"context,omitempty"`
	Decision     string    `json:"decision,omitempty"`
	Consequences string    `json:"consequences,omitempty"`
}

var appFs = afero.NewOsFs()

func (d DecisionRecord) format() (bytes.Buffer, error) {
	tmpl, err := afero.ReadFile(appFs, "templates/template.md")
	if err != nil {
		return bytes.Buffer{}, err
	}
	initial, err := template.New("initial").Parse(string(tmpl))
	if err != nil {
		return bytes.Buffer{}, err
	}

	var wr bytes.Buffer
	fmtwr := bufio.NewWriter(&wr)
	err = initial.Execute(fmtwr, d)
	if err != nil {
		return bytes.Buffer{}, err
	}
	return wr, nil
}

func (d DecisionRecord) formatTitle() string {
	underscorer := strings.NewReplacer(" ", "_")
	return fmt.Sprintf("%d-%s", d.Number, underscorer.Replace(d.Title))
}

func (d DecisionRecord) writeFile(path string) error {
	b, err := d.format()
	if err != nil {
		return err
	}
	p := filepath.Join(path, d.formatTitle())
	return afero.WriteFile(appFs, p, b.Bytes(), 0644)
}

// NextDecisionRecordNumber will get the next number in line for the ADR
func NextDecisionRecordNumber(path string) (int, error) {
	files, err := afero.ReadDir(appFs, path)
	if err != nil {
		return -1, err
	}
	highest := 0
	for _, f := range files {
		if !f.IsDir() {
			nm := f.Name()
			if n, err := parseADRNumber(nm); err == nil {
				if highest < n {
					highest = n
				}
			}
		}
	}
	return highest + 1, nil
}

func parseADRNumber(fname string) (int, error) {
	num := strings.Split(fname, "-")[0]
	return strconv.Atoi(num)
}
