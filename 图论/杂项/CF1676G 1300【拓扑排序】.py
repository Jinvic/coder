# import math
# from decimal import *
# 拓扑排序
t = int(input())
while(t > 0):
    t -= 1
    n = int(input())
    p = list(map(int, input().split(" ")))
    s = " "+input()  # 使下标和序号一致
    d = [0]*(n+1)  # 存入度
    a = [0]*(n+1)  # 存黑白节点差值
    b = []  # 存入度为0的节点用于遍历
    c = []  # 暂存入度变为0的节点，之后存入b
    for x in p:  # 计算入度
        d[x] += 1
    for i in range(1, n+1):  # 找出入度为0的节点
        if(d[i] == 0):
            b.append(i)
    p.insert(0, 0)  # 使下标和序号一致
    p.insert(1, 0)
    res = 0  # 结果
    cnt = 0  # 已统计过的点数量
    while(cnt < n):
        for x in b:
            if(s[x] == "B"):  # 点本身的颜色
                a[x] += 1
            else:
                a[x] -= 1
            if(a[x] == 0):  # 满足条件
                res += 1
            cnt += 1  # 计数
            a[p[x]] += a[x]  # 结果加到父节点
            d[p[x]] -= 1  # 删去这个点，父节点入度-1
            if(d[p[x]] == 0):  # 父节点成为新的叶子节点
                c.append(p[x])  # 存入c
        b = list(set(c))  # 转存b，用于下次统计
        c = []  # 清空c
    print(res)
