import random
from sympy import *
from sympy.solvers import solve

def find_factorial(n):
    if n == 0:
        return Symbol('f')(*[x**2 for x in range(1, (n+1))])
    elif n % 4 == 0: # n is a multiple of 4
        return Symbol('f')(*[((x-3) * (x-1) // 2 + y * (y-1) // 2) for x in range(1, n) if x != n and 0 < x <= 3], f=x+2)
    elif n % 4 == 1: # n is odd
        return Symbol('f')(*[(((x-1)**2 + (y-1)) // 2 - 1) for x in range(1, n) if x != n and 0 < x <= 3], f=x+2)
    elif n % 4 == 3: # n is a multiple of 8
        return Symbol('f')(*[(((x-1)**3 + (y-1)) // 2 - 1) for x in range(1, n)], f=0)

def find_divisors(n):
    divs = []
    i = 1
    while True:
        if n % i == 0:
            divs.append(i)
            n /= i
        else:
            break
    return divs

n_values = [4, 6, 8, 12, 18, 36]

factorials = []
for n in n_values:
    f = find_factorial(n)
    factorials.append(f)

divisors = []
for n in n_values:
    d = find_divisors(n)
    divisors.append(d)

print("Factorial values:", factorials)
print("Divisors:", divisors)
