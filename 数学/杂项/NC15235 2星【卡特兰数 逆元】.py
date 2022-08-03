# import math
# from decimal import *

m = 998244353
N = 200005


def quickpow(b, p):
    if (p < 0):
        return 0
    ret = 1
    b %= m
    while (p):
        if (p & 1):
            ret = (ret * b) % m
        p >>= 1
        b = (b * b) % m
    return ret


def inv(a):
    return quickpow(a, m - 2)


def cul(x):
    r = 1
    for i in range(1, x+1):
        r *= i
        r %= m
    return r


fac = [0]*N
facinv = [0]*N
if __name__ == '__main__':
    fac[0] = 1
    for i in range(1, N):
        fac[i] = fac[i-1]*i % m
    facinv[N-1] = inv(fac[N-1])
    for i in range(N-2, -1, -1):
        facinv[i] = facinv[i+1]*(i+1) % m
    n = int(input())
    for nn in range(n):
        x = int(input())
        if(x == 1):
            print(0)
        else:
            x -= 1
            r1 = fac[2*x]
            r2 = facinv[x+2]
            r3 = facinv[x-1]
            # print(r1, r2, r3)
            r = (3*r1) % m*r2 % m*r3 % m
            print("Case #"+str(nn+1)+": ", end="")
            print(int(r))
