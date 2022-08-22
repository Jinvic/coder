#https://codeforces.com/contest/1703/problem/E
# import math
# from decimal import *


t = int(input())
while(t > 0):
    t -= 1
    n = int(input())
    s = []
    for i in range(n):
        s.append(input())
    m = n//2
    res = 0
    for i in range(m):
        for j in range(m):
            sum = 0
            sum += int(s[i][j])
            sum += int(s[j][n-1-i])
            sum += int(s[n-1-i][n-1-j])
            sum += int(s[n-1-j][i])
            res += min(sum, 4-sum)
    if(n & 1 == 1 and n > 1):
        for i in range(m):
            sum = 0
            sum += int(s[i][m])
            sum += int(s[m][n-1-i])
            sum += int(s[n-1-i][m])
            sum += int(s[m][i])
            res += min(sum, 4-sum)
    print(res)
