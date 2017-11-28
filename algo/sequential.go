package algo

import "math/big"

// Factorize continuously runs Pollard's rho algorithm, until factor is found.
func Factorize(n *big.Int) *big.Int {
	// set initial values
	var (
		p0 = new(big.Int)
		c  = new(big.Int)
		f  = big.NewInt(0)
	)

	for f.Cmp(zero) == 0 {
		// generate random p0 and c
		globalRand.nextInt(p0, n)
		globalRand.nextInt(c, n)

		// try find factor
		f = rho(n, p0, c)
	}

	return f
}

// rho is Pollard's rho algorithm, with f(x) = x^2 + c, and x0 = p0.
// Returns "factor" found by Floyd's cycle detection algorithm.
// It can be 0(try again), or actual factor.
func rho(n, p0, c *big.Int) *big.Int {
	// set initial values
	var (
		x = new(big.Int).Set(p0)
		y = new(big.Int).Set(p0)
		g = big.NewInt(1)
	)

	for g.Cmp(one) == 0 {
		// x = (x*x + c) % n
		x.Mod(x.Add(x.Mul(x, x), c), n)

		// y = (y*y + c) % n
		y.Mod(y.Add(y.Mul(y, y), c), n)
		// once more
		y.Mod(y.Add(y.Mul(y, y), c), n)

		// g = gcd(abs(x - y), n)
		g.GCD(nil, nil, g.Abs(g.Sub(x, y)), n)
	}

	return g
}
