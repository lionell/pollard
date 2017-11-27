# Concurrent Rho algorithm

This is concurrent implementation of Pollard's Rho algorithm in Go. It also contains
sequential implementation in Python(see `main.py`).

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

Here are a couple of my **empirical** observations:

* There is number 4 that algorithm is failing to factorize.
* For small numbers(eg. 8, 16, 25) concurrent solution can find factor faster.
  These numbers have high probability of cycling in Rho procedure. Only some specific pairs `(x0, c)`
  lead to the result.
* For numbers(even very big) that have many small factors, algorithm will find answer very quick.
  Parallel execution here **will not help**.
* For numbers with only a few factors, concurrent algorithms can help.


Here are benchmark results from running program on my laptop(Thinkpad X1 Carbon 4Gen).

BenchmarkFactorizeParallel6\_1-4      |      10000    |      13240036 ns/op
BenchmarkFactorizeParallel6\_2-4      |      10000    |      13176139 ns/op
BenchmarkFactorizeParallel6\_3-4      |      10000    |      13681833 ns/op
BenchmarkFactorizeParallel6_4-4      |       5000    |      16099252 ns/op
BenchmarkFactorizeParallel7_1-4      |       2000    |      54945878 ns/op
BenchmarkFactorizeParallel7_2-4      |       2000    |      57374638 ns/op
BenchmarkFactorizeParallel7_3-4      |       2000    |      53736698 ns/op
BenchmarkFactorizeParallel7_4-4      |       2000    |      63402946 ns/op
BenchmarkFactorizeParallel8_1-4      |        500    |     201967658 ns/op
BenchmarkFactorizeParallel8_2-4      |        500    |     205176048 ns/op
BenchmarkFactorizeParallel8_3-4      |        500    |     194191468 ns/op
BenchmarkFactorizeParallel8_4-4      |        300    |     243370318 ns/op

