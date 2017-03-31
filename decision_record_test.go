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
	"bytes"
	"reflect"
	"testing"
)

func TestDecisionRecord_format(t *testing.T) {
	tests := []struct {
		name    string
		d       DecisionRecord
		want    bytes.Buffer
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.format()
			if (err != nil) != tt.wantErr {
				t.Errorf("DecisionRecord.format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecisionRecord.format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecisionRecord_formatTitle(t *testing.T) {
	tests := []struct {
		name string
		d    DecisionRecord
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.formatTitle(); got != tt.want {
				t.Errorf("DecisionRecord.formatTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecisionRecord_writeFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		d       DecisionRecord
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.d.writeFile(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("DecisionRecord.writeFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNextDecisionRecordNumber(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NextDecisionRecordNumber(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NextDecisionRecordNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NextDecisionRecordNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseADRNumber(t *testing.T) {
	type args struct {
		fname string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			"first test",
			args{"0002-Some Thoughts"},
			2,
			false,
		},
		{
			"second test, fail",
			args{"junkthing"},
			0,
			true,
		},
		{
			"third test, different ADR format",
			args{"0200 - lots of things.md"},
			0,
			true,
		},
		{
			"large num",
			args{"2000530-whoa.md"},
			2000530,
			false,
		},
		{
			"middle num",
			args{"25-thing.md"},
			25,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseADRNumber(tt.args.fname)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseADRNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseADRNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
