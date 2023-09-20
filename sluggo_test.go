package sluggo

import (
	"testing"
)

func TestSluggo(t *testing.T) {
	testTable := []struct {
		in  string
		out string
	}{
		{
			in:  "Hello",
			out: "hello",
		},
		{
			in:  "Hello 😊, 123",
			out: "hello-123",
		},
		{
			in:  "Money & Power = ?",
			out: "money-power",
		},
	}

	for _, tt := range testTable {
		t.Run(tt.in, func(t *testing.T) {
			s := Sluggo(tt.in)
			if s != tt.out {
				t.Errorf("got %s, want %s", s, tt.out)
			}
		})
	}
}
