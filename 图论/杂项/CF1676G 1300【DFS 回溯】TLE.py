#https://codeforces.com/contest/1676/problem/G
# import math
# from decimal import *

# 搜索函数
# 参数：当前节点 已搜索到黑白节点数相等的根结点数res
# 返回值：黑白节点差 res
def bw(u, res):
    if(len(g[u]) == 0):  # 为叶子节点
        if(s[u] == "B"):
            return 1, res
        else:
            return -1, res

    if(s[u] == "B"):  # 根结点颜色
        ret = 1
    else:
        ret = -1
    for v in g[u]:  # 遍历字树统计节点
        tmp, res = bw(v, res)
        ret += tmp
    if(ret == 0):  # 黑白节点数相等，res+1
        res += 1
    return ret, res


t = int(input())
while(t > 0):
    t -= 1
    n = int(input())
    p = list(map(int, input().split(" ")))
    g = [[] for i in range(n+1)]
    for i in range(0, n-1):
        g[p[i]].append(i+2)
    s = input()
    s = " "+s

    tmp, res = bw(1, 0)
    print(res)

