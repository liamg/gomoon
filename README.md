# gomoon
[![Build Status](https://travis-ci.org/liamg/gomoon.svg?branch=master)](https://travis-ci.org/liamg/gomoon)
[![GoDoc](https://godoc.org/github.com/liamg/gomoon?status.svg)](https://godoc.org/github.com/liamg/gomoon)

Need to fire a lambda on each full moon?

Typical business requirements of generating some client reports every [waxing gibbous](https://en.wikipedia.org/wiki/Waxing_gibbous)? 

No problem! This is a package for calculating the approximate<sup>[1](#footnote1)</sup> moon phase for all of your lunar needs.

## Usage

```go
package main

import (
	"fmt"

	"github.com/liamg/gomoon"
)

func main() {
	if gomoon.PhaseNow() == gomoon.FULL_MOON {
		fmt.Println("It's a full moon!")
	}
}
```

---

<a name="footnote1">1</a>: This is a dubious claim at best, moon phase calculations are half assed and may let you down if you require anything more than "vague" accuracy.
