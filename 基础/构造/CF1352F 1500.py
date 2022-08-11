#https://codeforces.com/contest/1352/problem/F
# import math
# from decimal import *

t = int(input())
while(t > 0):
    t -= 1
    a, b, c = map(int, input().split(" "))
    if(b == 0):
        if(a != 0):
            print("0"*(a+1))
        else:
            print("1"*(c+1))
    else:
        if(b & 1):
            sl = ""
        else:
            b -= 1
            sl = "1"
        d = b//2+1
        s01 = "01"*d
        s0 = "0"*a
        s1 = "1"*c
        print(sl+s0+s01+s1)
