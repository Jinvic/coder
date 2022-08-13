#https://codeforces.com/contest/1676/problem/F
# import math
# from decimal import *

t = int(input())
while(t > 0):
    t -= 1
    n, k = map(int, input().split(" "))
    a = list(map(int, input().split(" ")))
    a.sort()
    b = []  # 存原数据
    c = []  # 存每个数的数量
    for i in range(n):
        if(i == 0):
            b.append(a[0])
            c.append(1)
        else:
            if(a[i] == a[i-1]):
                c[-1] += 1
            else:
                b.append(a[i])
                c.append(1)
    ll = 0
    rr = -1
    flag = False  # 是否有解
    for l in range(len(c)):
        if(c[l] >= k):  # 确定左端点
            flag = True
            fflag = True  # 是否到最右
            for r in range(l+1, len(c)):  # 确定右端点
                if(c[r] < k or b[r] != b[r-1]+1):
                    r -= 1
                    if(r-l > rr-ll):  # 更新最优解
                        ll = l
                        rr = r
                    fflag = False
                    l = r
                    break
            if(fflag):  # 右端点为最右
                r = len(c)-1
                if(r-l > rr-ll):
                    ll = l
                    rr = r
                break
    if(flag):
        print(b[ll], b[rr])
    else:
        print(-1)
