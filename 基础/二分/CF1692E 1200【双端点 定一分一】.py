# https://codeforces.com/contest/1692/problem/E
# import math
# from decimal import *


t = int(input())
while(t > 0):
    t -= 1
    n, s = map(int, input().split(" "))
    a = list(map(int, input().split(" ")))
    pre = [0]*n
    sum = a[0]
    pre[0] = a[0]
    for i in range(1, n):
        pre[i] = a[i]+pre[i-1]
        sum += a[i]
    if(sum < s):
        print(-1)
        continue
    res = n
    for i in range(n):
        if(i == 0):
            lp = 0
        else:
            lp = pre[i-1]
        if(sum-lp < s):
            break
        l = i
        r = n-1
        ans = n
        while(l <= r):
            mid = (l+r) >> 1
            if(pre[mid]-lp > s):
                r = mid-1
            else:
                if(pre[mid]-lp == s):
                    ans = mid
                l = mid+1
        res = min(res, n-(ans-i+1))
    print(res)
