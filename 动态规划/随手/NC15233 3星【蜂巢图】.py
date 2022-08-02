# import math
# from decimal import *


if __name__ == '__main__':
    n = int(input())
    h = 4*n-3
    d = n-1
    l = 2*n-1
    dp = [0]*(2*n+5)
    dp[n] = int(input())
    for i in range(1, d):
        k = n-i
        lst = list(map(int, input().split(" ")))
        for x in lst:
            if n-i == k:  # left side
                dp[k] = x+dp[k+1]
            elif n+i == k:
                dp[k] = x+dp[k-1]
            else:
                dp[k] = x+max(dp[k], dp[k-1], dp[k+1])
            k += 2
    for i in range(d, h-d):
        if (i-d) & 1:
            k = 2
        else:
            k = 1
        lst = list(map(int, input().split(" ")))
        for x in lst:
            if k == 1:
                if(i == d):
                    dp[k] = x+dp[k+1]
                else:
                    dp[k] = x+max(dp[k], dp[k+1])
            elif k == l:
                if(i == d):
                    dp[k] = x+dp[k-1]
                else:
                    dp[k] = x+max(dp[k], dp[k-1])
            else:
                dp[k] = x+max(dp[k], dp[k-1], dp[k+1])
            k += 2
    for i in range(d-1, -1, -1):
        k = n-i
        lst = list(map(int, input().split(" ")))
        for x in lst:
            dp[k] = x + max(dp[k], max(dp[k - 1], dp[k + 1]))
            k += 2
    print(dp[n])
