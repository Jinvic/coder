//https://codeforces.com/problemset/problem/1354/B
#include <bits/stdc++.h>
using namespace std;
const int inf = 0x7f7f7f7f;
const int N = 200005;
int dp[N][4];//分别存最近的123的位置
int cul(int *d)
{
    int ma = max(max(d[1], d[2]), d[3]),
        mi = min(min(d[1], d[2]), d[3]);
    if (mi == 0)
        return inf;
    else
        return ma - mi + 1;
}
int main(void)
{
    ios::sync_with_stdio(false);
    cin.tie(0), cout.tie(0);
    int t;
    string s;
    cin >> t;
    while (t--)
    {
        cin >> s;
        s = " " + s;
        for (int i = 1; i <= 4; i++)
            dp[0][i] = 0;
        int res = inf;
        for (int i = 1; i < s.size(); i++)
        {
            int c = s[i] - '0';
            for (int j = 1; j <= 3; j++)
            {
                if (j == c)
                    dp[i][j] = i;
                else
                    dp[i][j] = dp[i - 1][j];
            }
            res = min(res, cul(dp[i]));
        }
        // cout << inf << endl;
        if (res == inf)
            res = 0;
        cout << res << endl;
    }
    return 0;
}
