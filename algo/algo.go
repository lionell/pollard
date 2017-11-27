package algo

import (
	"math/big"
	"math/rand"
	"sync"
)

// Thread-safe version of rand.Rand
type lockedRand struct {
	r *rand.Rand
	sync.Mutex
}

func (lr *lockedRand) nextInt(z, n *big.Int) {
	lr.Lock()
	z.Rand(lr.r, n)
	lr.Unlock()
}

var (
	zero       = big.NewInt(0)
	one        = big.NewInt(1)
	globalRand = &lockedRand{r: rand.New(rand.NewSource(17))}
)
