package gomoon

/*
o                     __...__     *
              *   .--'    __.=-.             o
     |          ./     .-'
    -O-        /      /
     |        /    '"/               *
             |     (+)
            |        \                         .
            |         \
 *          |   |   ___\                  |
             |  .   /  `                 -O-
              \  `~~\                     |
         o     \     \            *
                `\    `-.__           .
    .             `--._    `--'jgs
                       `---~~`                *
            *                   o
*/

import (
	"math"
	"time"
)

type MoonPhase uint8

const (
	NEW MoonPhase = iota
	WAXING_CRESCENT
	FIRST_QUARTER
	WAXING_GIBBOUS
	FULL_MOON
	WANING_GIBBOUS
	LAST_QUARTER
	WANING_CRESCENT
	phaseCount
)

const (
	lunationDays     = 29.530588
	knownNewMoonUnix = 592500 // 1970-01-07T20:35:00Z
	secondsPerDay    = 86400
)

// Phase returns the current moon phase using the provided constants, so you can do, for example `if gomoon.PhaseNow() == gomoon.FULL_MOON { ... `
func PhaseNow() MoonPhase {
	return Phase(time.Now())
}

// Phase returns the moon phase using the provided constants, so you can do, for example `if gomoon.Phase(time.Now()) == gomoon.FULL_MOON { ... `
func Phase(t time.Time) MoonPhase {
	phase := calculatePhase(t)
	spread := (phase * float64(phaseCount)) + 0.5
	if spread >= float64(phaseCount) {
		return NEW
	}
	for i := phaseCount - 1; i >= 0; i-- {
		if spread > float64(i) {
			return MoonPhase(i)
		}
	}
	return NEW
}

// returns moon phase expressed in [0, 1], where 0 and 1 both represent new moon, 0.5 represents full moon
func calculatePhase(t time.Time) float64 {
	knownNewMoonTime := time.Unix(knownNewMoonUnix, 0)
	phase := (t.UTC().Sub(knownNewMoonTime).Seconds() / float64(secondsPerDay)) / lunationDays
	return phase - math.Floor(phase)
}
