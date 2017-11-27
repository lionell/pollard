#!/usr/bin/env python3

from fractions import gcd
import random

random.seed(17)

def pollard_rho(n):
    while True:
        c = random.randint(2, n-1)
        f = lambda x: x**2 + c 
        x = y = 2 
        d = 1 
        while d == 1:
            x = f(x) % n 
            y = f(f(y)) % n 
            d = gcd(abs(x - y), n)
        if d != n: return d

factor = pollard_rho(int(input()))
print("Factor {} found.".format(factor))
