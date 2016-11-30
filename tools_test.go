package adr

import (
	"testing"
	"time"
)

func TestDecisionRecord_format(t *testing.T) {
	type fields struct {
		Number       int
		Title        string
		Date         time.Time
		Status       string
		Context      string
		Decision     string
		Consequences string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := DecisionRecord{
				Number:       tt.fields.Number,
				Title:        tt.fields.Title,
				Date:         tt.fields.Date,
				Status:       tt.fields.Status,
				Context:      tt.fields.Context,
				Decision:     tt.fields.Decision,
				Consequences: tt.fields.Consequences,
			}
			got, err := d.format()
			if (err != nil) != tt.wantErr {
				t.Errorf("DecisionRecord.format() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecisionRecord.format() = %v, want %v", got, tt.want)
			}
		})
	}
}
