package main

import (
	"flag"
	"fmt"

	"github.com/kalebhawkins/goflipper/flipper"
)

func main() {
	flips := flag.Int("count", 100, "number of times to flip the coin")
	verbose := flag.Bool("verbose", false, "displays how many times the coin rotates during each flip")
	flag.Parse()

	c := flipper.Coin{true, false, 0, 0}

	for i := 0; i <= *flips; i++ {
		flipCount := c.Flip()
		if *verbose {
			fmt.Printf("Coin rotated %d times during this flip.\n", flipCount)
		}
	}

	fmt.Printf("The coin has been flipped %d times.\n", *flips)
	fmt.Println(c)
}
