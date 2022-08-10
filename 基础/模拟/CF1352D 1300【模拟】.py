#https://codeforces.com/contest/1352/problem/D
# import math
# from decimal import *

t = int(input())
while(t > 0):
    t -= 1
    n = int(input())
    lst = list(map(int, input().split(" ")))
    step = 1
    pa = lst[0]
    pb = 0
    suma = lst[0]
    sumb = 0
    del(lst[0])
    aturn = False
    while(len(lst) > 0):
        step += 1
        if(aturn):
            tmp = 0
            while(tmp <= pb and len(lst) > 0):
                tmp += lst[0]
                del(lst[0])
            if(tmp > pb):
                pa = tmp
            suma += tmp
        else:
            tmp = 0
            while(tmp <= pa and len(lst) > 0):
                tmp += lst[-1]
                del(lst[-1])
            if(tmp > pa):
                pb = tmp
            sumb += tmp
        aturn = not(aturn)
    print(step, suma,  sumb)
