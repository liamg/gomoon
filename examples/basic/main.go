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
