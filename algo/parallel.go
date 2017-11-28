package algo

import (
	"log"
	"math/big"
)

// FactorizeParallel continuously runs Pollard's rho algorithm, until factor is found.
// It runs j goroutines trying to find factor using different p0 and c.
// When the factor is found, we signalize other goroutines to stop and wait for them.
func FactorizeParallel(n *big.Int, j int) *big.Int {
	// set initial values
	var (
		res    = make(chan *big.Int, j)
		stop   = make(chan struct{}, j)
		factor = new(big.Int)
	)

	// start j goroutines
	for i := 0; i < j; i++ {
		startRho(n, res, stop)
	}

	for {
		factor = <-res
		if factor.Cmp(zero) != 0 {
			break
		}
		startRho(n, res, stop)
	}

	// signalize other goroutines
	for i := 0; i < j-1; i++ {
		stop <- struct{}{}
	}

	// wait for other goroutines to finish
	for i := 0; i < j-1; i++ {
		<-res
	}
	return factor
}

func startRho(n *big.Int, res chan<- *big.Int, stop <-chan struct{}) {
	var (
		p0 = new(big.Int)
		c  = new(big.Int)
	)
	globalRand.nextInt(p0, n)
	globalRand.nextInt(c, n)

	log.Printf("Running routine(n = %v, p0 = %v, c = %v).", n, p0, c)
	go parallelRho(n, p0, c, res, stop)
}

// parallelRho is Pollard's rho algorithm, with f(x) = x^2 + c, and x0 = p0.
// Returns "factor" found by Floyd's cycle detection algorithm.
// It can be 0(try again), or actual factor.
func parallelRho(n, p0, c *big.Int, res chan<- *big.Int, stop <-chan struct{}) {
	// set initial values
	var (
		x = new(big.Int).Set(p0)
		y = new(big.Int).Set(p0)
		g = big.NewInt(1)
		i = 0
	)

	for g.Cmp(one) == 0 {
		select {
		case <-stop:
			res <- big.NewInt(-1)
			return
		default:
			// x = (x*x + c) % n
			x.Mod(x.Add(x.Mul(x, x), c), n)

			// y = (y*y + c) % n
			y.Mod(y.Add(y.Mul(y, y), c), n)
			// once more
			y.Mod(y.Add(y.Mul(y, y), c), n)

			// g = gcd(abs(x - y), n)
			g.GCD(nil, nil, g.Abs(g.Sub(x, y)), n)

			i++
		}
	}

	log.Printf("%v iterations. (%v, %v)", i, p0, c)
	res <- g
}
