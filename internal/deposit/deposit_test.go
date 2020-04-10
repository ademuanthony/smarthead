package deposit

import (
	"testing"
	"time"
)

func TestNextMonday(t *testing.T)  {
	now := time.Now()
	for i := 0; i <= 7; i++ {
		d := now.Add(time.Duration(i) * 24 * time.Hour)
		date := NextMonday(d)
		if date.Weekday() != time.Monday {
			t.Logf("Expected %v got %v", time.Monday.String(), date.Weekday().String())
			t.Fail()
		}

		if date.Unix() > d.Add(7 * 24 * time.Hour).Unix() {
			t.Errorf("Expected the next monday to be within one week")
			t.Fail()
		}
	}
}