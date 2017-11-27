#!/usr/bin/env python3

'''
Generate big n-digit number
'''
import random

s = str(random.randint(1, 9))
for i in range(int(input()) - 1):
    s += str(random.randint(0, 9))

print(s)
