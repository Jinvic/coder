#https://codeforces.com/contest/1676/problem/H2
# import math
# from decimal import *

def lowbit(x):
    """
    x 的二进制表示中，最低位的 1 的位置。
    lowbit(0b10110000) == 0b00010000
             ~~~^~~~~
    lowbit(0b11100100) == 0b00000100
             ~~~~~^~~
    """
    return x & -x


def add(x, k):
    while x <= n:  # 不能越界
        c[x] = c[x] + k
        x = x + lowbit(x)


def getsum(x):  # a[1]..a[x]的和
    ans = 0
    while x >= 1:
        ans = ans + c[x]
        x = x - lowbit(x)
    return ans


t = int(input())
while(t > 0):
    t -= 1
    n = int(input())
    a = list(map(int, input().split(" ")))
    num = [0]*(n+1)
    c = [0]*(n+1)
    res = 0
    for i in range(n):
        num[a[i]] += 1
        res += i-getsum(a[i])  # 统计逆序对数量
        add(a[i], 1)
    for i in range(1, n+1):
        res += int(num[i]*(num[i]-1)/2)  # 统计同一区间交点数量
    print(res)
