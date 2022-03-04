// 2020ICPC上海站 B. Mine Sweeper II
// https://codeforces.com/gym/102900/problem/B
#include <iostream>
#include <cstdio>
#include <cstring>
#include <algorithm>
using namespace std;
const int maxn = 1003;
char a[maxn][maxn], b[maxn][maxn];
char revc(char c)
{
    if (c == 'X')
        return '.';
    else
        return 'X';
}
int main(void)
{
    int n, m, num = 0;
    scanf("%d%d", &n, &m);
    for (int i = 0; i < n; i++)
        scanf("%s", a[i]);
    for (int i = 0; i < n; i++)
        scanf("%s", b[i]);
    for (int i = 0; i < n; i++)
        for (int j = 0; j < m; j++)
            if (a[i][j] == b[i][j])
                num++;
    if (num > n * m / 2)
    {
        for (int i = 0; i < n; i++)
            printf("%s\n", a[i]);
    }
    else
    {
        for (int i = 0; i < n; i++)
        {
            for (int j = 0; j < m; j++)
                printf("%c", revc(a[i][j]));
            puts("");
        }
    }
    return 0;
}
