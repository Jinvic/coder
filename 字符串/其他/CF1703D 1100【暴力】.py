#https://codeforces.com/contest/1703/problem/D
# import math
# from decimal import *


t = int(input())
while(t > 0):
    t -= 1
    n = int(input())
    s = []
    d = {}
    for i in range(n):
        s.append(input())
        d[s[-1]] = 1
    res = ["0"]*n
    for i in range(n):
        l = len(s[i])
        if(l > 1):
            for j in range(1, l):
                if(s[i][:j] in d and s[i][j:] in d):
                    res[i] = "1"
                    break
    print("".join(res))
