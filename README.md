# Concurrent Rho algorithm

This is concurrent implementation of Pollard's Rho algorithm in Go. It also contains
sequential implementation in Go and Python(see main.py).

## How it works

Rho algorithm is generating the sequence of numbers `x0, x1 = f(x0), ..., x_i = f(x_{i-1})`,
where `f(x) = x*x + c`.
We are applying Floyd's cycle detection algorithm, to find point in the cycle and this can give 
us factor. For more details see [wiki](https://en.wikipedia.org/wiki/Pollard%27s_rho_algorithm).

We make it concurrent by running with different starting parameters `x0` and `c`. As soon as one of the
subroutines find factor, we stop all the other. This simple **concurrent** modification of algorithm can
give us performance gain.

## Results

As you know, Rho algorithm belongs to randomized, so it's hard to get accurate performance stats.

Here is a couple of **empirical** results I've got from testing this algorithm:
* There is number 4 that algorithm is failing to factorize.
* For small numbers(eg. 8, 16, 25) concurrent solution can find factor faster.
  These numbers have high probability of cycling in Rho procedure. Only some specific pairs (p0, c)
  lead to the result.
* For numbers(even very big) that have many small factors, algorithm will find answer very quick.
  Parallel execution here **will not help**.
* For numbers with only a few factors, concurrent algorithms can help.
  Here are benchmark results from running program on my laptop(Thinkpad X1 Carbon 4Gen).

