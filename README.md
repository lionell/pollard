# Concurrent Rho algorithm

[![Build Status](https://travis-ci.org/lionell/pollard.svg?branch=master)](https://travis-ci.org/lionell/pollard)

Implementation of Pollard's Rho algorithm in Go. It also contains
sequential implementation in Python(see `main.py`).

## How it works

Rho algorithm is generating the sequence of numbers `x0, x1 = f(x0), ..., x_i = f(x_{i-1})`,
where `f(x) = x*x + c`.
We are applying Floyd's cycle detection algorithm, to find point in the cycle and this can give 
us factor. For more details see [wiki](https://en.wikipedia.org/wiki/Pollard%27s_rho_algorithm).

We make it concurrent by running with different starting parameters `x0` and `c`. As soon as one of the
subroutines find factor, we stop all the other. This simple **concurrent** modification of algorithm can
give us performance gain.

## Usage

To run sequential version simply call `algo.Factorize(n)`. If you want to use concurrent implementation,
you'll need to specify additional parameter `j` which is concurrency limit. Eg. `algo.FactorizeParallel(n, 4)`.

To run benchmarks just do:
```(shell)
$ go test -bench . -benchtime 1m -timeout 1h ./...
```

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


Here are benchmark results from running program on my laptop(Intel Core i7-6600U @ 4x 3.4GHz).

| Name               |       Runs | Time(ms/op) |
|:-------------------|-----------:|------------:|
| Parallel\_6\_1     |      10000 |      13.240 |
| Parallel\_6\_2     |      10000 |      13.176 |
| Parallel\_6\_3     |      10000 |      13.681 |
| Parallel\_6\_4     |       5000 |      16.099 |
| Parallel\_7\_1     |       2000 |      54.945 |
| Parallel\_7\_2     |       2000 |      57.374 |
| Parallel\_7\_3     |       2000 |      53.736 |
| Parallel\_7\_4     |       2000 |      63.402 |
| Parallel\_8\_1     |        500 |     201.967 |
| Parallel\_8\_2     |        500 |     205.176 |
| Parallel\_8\_3     |        500 |     194.191 |
| Parallel\_8\_4     |        300 |     243.370 |
| Parallel\_9\_1     |         50 |    2567.833 |
| Parallel\_9\_2     |         50 |    2264.897 |
| Parallel\_9\_3     |         30 |    2317.177 |
| Parallel\_9\_4     |         50 |    2005.279 |

Where each benchmark is based on `prod prime[i] for i from left to right`.

| Name            |       Left |      Right |
|:----------------|-----------:|-----------:|
| Parallel\_6     |    1000000 |    1000005 |
| Parallel\_7     |   10000000 |   10000005 |
| Parallel\_8     |  100000000 |  100000005 |
| Parallel\_9     | 1000000001 | 1000000005 |

And last part of benchmark name states for concurrency limit set for the test.
