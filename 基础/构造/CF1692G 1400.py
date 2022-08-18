# https://codeforces.com/contest/1692/problem/G
# import math
# from decimal import *


t = int(input())
while(t > 0):
    t -= 1
    n, k = map(int, input().split(" "))
    a = list(map(int, input().split(" ")))
    b = []
    for i in range(n-1):
        if(a[i] < a[i+1]*2):
            b.append(1)
        else:
            b.append(0)
    b.insert(0, 0)
    for i in range(1, n):
        b[i] += b[i-1]
    cnt = 0
    for i in range(n-k):
        if(b[i+k]-b[i] == k):
            cnt += 1
    print(cnt)
