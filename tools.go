package adr

import "text/template"
import "time"
import "io/ioutil"

import "bufio"
import "bytes"

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

func (d DecisionRecord) format() (string, error) {
	tmpl, err := ioutil.ReadFile("templates/template.md")
	if err != nil {
		return "", err
	}
	initial, err := template.New("initial").Parse(string(tmpl))
	if err != nil {
		return "", err
	}

	var wr bytes.Buffer
	fmtwr := bufio.NewWriter(&wr)
	err = initial.Execute(fmtwr, d)
	if err != nil {
		return "", err
	}
	return wr.String(), nil
}

func (d DecisionRecord) writeFile(path string) error {

}
