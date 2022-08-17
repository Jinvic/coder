#https://codeforces.com/contest/1676/problem/H1
# import math
# from decimal import *

t = int(input())
while(t > 0):
    t -= 1
    n = int(input())
    a = list(map(int, input().split(" ")))
    num = [0]*(n+1)
    res = 0
    for i in range(n):
        sum = 0
        for j in range(1, a[i]+1):  # 统计逆序对数量
            sum += num[j]
        res += i-sum
        num[a[i]] += 1
    for i in range(1, n+1):  # 统计同一区间交点数量
        res += int(num[i]*(num[i]-1)/2)
    print(res)
