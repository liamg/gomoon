package gomoon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestKnownMoonPhases(t *testing.T) {

	tests := []struct {
		date  time.Time
		phase MoonPhase
	}{
		{
			date:  mustParseTime("1900-01-01T12:00:00Z"),
			phase: NEW,
		},
		{
			date:  mustParseTime("2000-01-06T19:24:01Z"),
			phase: NEW,
		},
		{
			date:  mustParseTime("1980-07-20T05:51:00Z"),
			phase: FIRST_QUARTER,
		},
		{
			date:  mustParseTime("1990-09-05T01:46:00Z"),
			phase: FULL_MOON,
		},
		{
			date:  mustParseTime("2000-07-24T11:03:00Z"),
			phase: LAST_QUARTER,
		},
	}

	for _, test := range tests {
		t.Run(test.date.String(), func(t *testing.T) {
			assert.Equal(t, test.phase, Phase(test.date))
		})
	}

}

func mustParseTime(t string) time.Time {
	parsed, err := time.Parse(time.RFC3339, t)
	if err != nil {
		panic(err)
	}
	return parsed
}
