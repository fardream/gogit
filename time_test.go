package gogit_test

import (
	"testing"
	"time"

	"github.com/fardream/gogit"
)

func TestTimeSecsOffset(t *testing.T) {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Logf("cannot load new york time zone: %s", err.Error())
		return
	}

	atime := time.Date(2023, time.September, 17, 20, 10, 9, 0, location)

	secs, offset := gogit.TimeSecsOffset(atime)
	if offset != "-0400" {
		t.Fatalf("offset is not -0400: %s", offset)
	}
	if secs != 1694981409 {
		t.Fatalf("time is not 1694981409 %d", secs)
	}
}
