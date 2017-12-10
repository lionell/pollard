from fractions import gcd
import random

random.seed(17)

def pollard_rho(n):
    i = 0
    while True:
        c = random.randint(2, n-1)
        f = lambda x: x**2 + c 
        x = y = 2 
        d = 1 
        while d == 1:
            x = f(x) % n 
            y = f(f(y)) % n 
            d = gcd(abs(x - y), n)
            i += 1
        if d != n: return d, i

factor, cnt = pollard_rho(int(input()))
print("Factor {} found in {} iterations.".format(factor, cnt))
