package adr

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// DecisionRecord is a type of Architecture Decision Record.
type DecisionRecord struct {
	Number       int
	Title        string
	Date         time.Time
	Status       string
	Context      string
	Decision     string
	Consequences string
}

func (d DecisionRecord) format() (bytes.Buffer, error) {
	tmpl, err := ioutil.ReadFile("templates/template.md")
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
	return ioutil.WriteFile(p, b.Bytes(), 0644)
}

// NextDecisionRecordNumber will get the next number in line for the ADR
func NextDecisionRecordNumber(path string) (int, error) {
	files, err := ioutil.ReadDir(path)
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
