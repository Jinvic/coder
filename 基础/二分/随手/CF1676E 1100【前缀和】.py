#https://codeforces.com/contest/1676/problem/E
# import math
# from decimal import *

t = int(input())
while(t > 0):
    t -= 1
    n, m = map(int, input().split(" "))
    a = list(map(int, input().split(" ")))
    a.sort(reverse=True)
    b = [0]*(len(a))
    b[0] = a[0]
    for i in range(1, len(a)):
        b[i] = a[i]+b[i-1]
    while(m > 0):
        m -= 1
        q = int(input())
        if(b[-1] < q):
            print(-1)
        else:
            l = 0
            r = len(b)-1
            ans = r
            while(l <= r):
                mid = (l+r) >> 1
                if(b[mid] >= q):
                    ans = mid
                    r = mid-1
                else:
                    l = mid+1
            print(ans+1)
