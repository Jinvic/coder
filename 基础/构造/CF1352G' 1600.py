#https://codeforces.com/contest/1352/problem/G
# import math
# from decimal import *

t = int(input())
while(t > 0):
    t -= 1
    n = int(input())
    if(n < 4):
        print(-1)
    elif(n & 1):
        a = []
        b = 1
        while(b <= n):
            a.append(b)
            b += 2
        a.append(n-3)
        a.append(n-1)
        c = n-5
        while(c >= 2):
            a.append(c)
            c -= 2
        print(" ".join(map(str, a)))
    else:
        a = []
        b = 1
        while(b < n-3):
            a.append(b)
            b += 2
        a.append(n-1)
        a.append(n-3)
        c = n
        while(c >= 2):
            a.append(c)
            c -= 2
        print(" ".join(map(str, a)))
