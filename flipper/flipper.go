package flipper

import (
	"fmt"
	"math/rand"
	"time"
)

type Coin struct {
	Heads  bool
	Tails  bool
	HCount int
	TCount int
}

func (c *Coin) Flip() (flipCount int) {
	seed := rand.NewSource(time.Now().UnixNano())
	randEngine := rand.New(seed)

	flipCount = randEngine.Intn(20) + 1
	c.Heads = true
	c.Tails = false

	for i := 0; i <= flipCount; i++ {
		c.Heads, c.Tails = c.Tails, c.Heads
	}

	if c.Heads {
		c.HCount += 1
	} else {
		c.TCount += 1
	}

	return
}

func (c Coin) String() string {
	return fmt.Sprintf("Heads: %t\nTails: %t\nHeads Count: %d\nTailsCount: %d", c.Heads, c.Tails, c.HCount, c.TCount)
}
