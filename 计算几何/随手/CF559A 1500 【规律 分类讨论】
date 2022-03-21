//https://codeforces.com/problemset/problem/559/A
#include <bits/stdc++.h>
using namespace std;
int tri[2005];

int main(void)
{
    tri[0] = 0;
    for (int i = 1; i < 2005; i++)
    {
        tri[i] = tri[i - 1] + 2 * i - 1;
    }
    int a[6];
    for (int i = 0; i < 6; i++)
    {
        cin >> a[i];
    }
    int sum = 0;
    if (a[1] == a[5])
    {
        int b = a[0] + a[1];
        sum += tri[b] - tri[a[0]];
        sum += tri[b] - tri[a[3]];
    }
    else
    {
        int c = min(a[1], a[5]);
        int b = a[0] + c;
        int d = abs(a[1] - a[5]);
        sum += tri[b] - tri[a[0]];
        sum += 2 * b * d;
        sum += tri[b] - tri[a[3]];
    }
    cout << sum;
}
