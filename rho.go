package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/lionell/pollard/algo"
	"math/big"
	"os"
)

var concurrencyLimit = flag.Int("j", 1, "max number of sub-routines to run concurrently")

func main() {
	flag.Parse()

	// read input
	s := bufio.NewScanner(os.Stdin)
	s.Scan()

	// try parse number
	n, ok := new(big.Int).SetString(s.Text(), 10)
	if !ok {
		fmt.Printf("Error while parsing input number %v", s.Text())
		return
	}

	// Check for primality. 10 iterations of Miller-Rabin + Baillie-PSW
	// is enough to say with certainty 1 - (1/4)^10
	if n.ProbablyPrime(10) {
		fmt.Printf("%v is prime.\n", n)
		return
	}

	// finally, try to factorize
	f := algo.FactorizeParallel(n, *concurrencyLimit)
	fmt.Printf("Factor %v found.\n", f)
}
