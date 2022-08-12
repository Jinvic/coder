//https://codeforces.com/contest/1406/problem/B

//dp，写假了，虽然ac
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 200005;
const int inf = 0x7f7f7f7f;
const int mod = 1000000007;
const double eps = 1e-8;

ll dp[N][6][2][2]; //到第及个数 取了几个数 结果正负 最大最小值
signed main(void)
{
    ios::sync_with_stdio(false);
    cin.tie(0), cout.tie(0);
    int t, n;
    ll x;
    cin >> t;
    while (t--)
    {
        cin >> n;
        memset(dp, 0, (sizeof(ll) * 6 * 2 * 2) * (n + 2));
        int cnt0 = 0;
        for (int i = 1; i <= n; i++)
        {
            cin >> x;
            if (x < 0)
            {
                for (int j = 1; j <= i && j <= 5; j++)
                {
                    for (int k = 0; k < 2; k++)
                        for (int l = 0; l < 2; l++)
                        {   
                            bool f1 = dp[i - 1][j][k][l], f2 = dp[i - 1][j - 1][!k][!l];
                            if (f1 && f2)
                            {
                                if (l)
                                    dp[i][j][k][l] = max(dp[i - 1][j][k][l], x * dp[i - 1][j - 1][!k][!l]);
                                else
                                    dp[i][j][k][l] = min(dp[i - 1][j][k][l], x * dp[i - 1][j - 1][!k][!l]);
                            }
                            else if (f1 && !f2)
                            {
                                dp[i][j][k][l] = dp[i - 1][j][k][l];
                                if (j == 1)
                                {
                                    if (l)
                                        dp[i][j][k][l] = max(dp[i][j][k][l], x);
                                    else
                                        dp[i][j][k][l] = min(dp[i][j][k][l], x);
                                }
                            }
                            else if (!f1 && f2)
                            {
                                dp[i][j][k][l] = x * dp[i - 1][j - 1][!k][!l];
                            }
                            else if (!f1 && !f2)
                            {
                                if (j == 1 && k == 0)
                                {
                                    dp[i][j][k][0] = x;
                                    dp[i][j][k][1] = x;
                                }
                            }
                        }
                }
            }
            else if (x > 0)
            {
                for (int j = 1; j <= i && j <= 5; j++)
                {
                    for (int k = 0; k < 2; k++)
                        for (int l = 0; l < 2; l++)
                        {
                            bool f1 = dp[i - 1][j][k][l], f2 = dp[i - 1][j - 1][k][l];
                            if (f1 && f2)
                            {
                                if (l)
                                    dp[i][j][k][l] = max(dp[i - 1][j][k][l], x * dp[i - 1][j - 1][k][l]);
                                else
                                    dp[i][j][k][l] = min(dp[i - 1][j][k][l], x * dp[i - 1][j - 1][k][l]);
                            }
                            else if (f1 && !f2)
                            {
                                dp[i][j][k][l] = dp[i - 1][j][k][l];
                                if (j == 1)
                                {
                                    if (l)
                                        dp[i][j][k][l] = max(dp[i][j][k][l], x);
                                    else
                                        dp[i][j][k][l] = min(dp[i][j][k][l], x);
                                }
                            }
                            else if (!f1 && f2)
                            {
                                dp[i][j][k][l] = x * dp[i - 1][j - 1][k][l];
                            }
                            else if (!f1 && !f2)
                            {
                                if (j == 1 && k == 1)
                                {
                                    dp[i][j][k][0] = x;
                                    dp[i][j][k][1] = x;
                                }
                            }
                        }
                }
            }
            else
            {
                cnt0++;
                for (int j = 1; j <= i && j <= 5; j++)
                {
                    for (int k = 0; k < 2; k++)
                        for (int l = 0; l < 2; l++)
                            dp[i][j][k][l] = dp[i - 1][j][k][l];
                }
            }
        }
        ll r = 0;
        if (dp[n][5][0][1] != 0) //存在负数最大解
            r = dp[n][5][0][1];
        if (dp[n][5][1][1] != 0) //存在正数最大解
            r = dp[n][5][1][1];
        if (r < 0 && cnt0)
            cout << 0 << endl;
        else
            cout << r << endl;
    }
    return 0;
}

//正解
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 200005;
const int inf = 0x7f7f7f7f;
const int mod = 1000000007;
const double eps = 1e-8;

ll a[N], b[N];
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	int t, n;
	ll x;
	cin >> t;
	while (t--)
	{
		cin >> n;
		int na = 0, nb = 0;
		for (int i = 0; i < n; i++)
		{

			cin >> x;
			if (x > 0)
				b[nb++] = x;
			else
				a[na++] = x;
		}
		sort(a, a + na);
		sort(b, b + nb);
		ll r = 0;
		if (nb == 0)
		{
			r = 1;

			for (int i = 1; i <= 5; i++)
				r *= a[na - i];
			cout << r << endl;
			continue;
		}
		if (nb >= 5)
		{
			ll tmp = 1;
			for (int i = 1; i <= 5; i++)
				tmp *= b[nb - i];
			r = max(r, tmp);
		}
		if (nb >= 3 && na >= 2)
		{
			ll tmp = 1;
			for (int i = 1; i <= 3; i++)
				tmp *= b[nb - i];
			tmp *= a[0] * a[1];
			r = max(r, tmp);
		}
		if (nb >= 1 && na >= 4)
		{
			ll tmp = 1;
			for (int i = 0; i < 4; i++)
				tmp *= a[i];
			tmp *= b[nb - 1];
			r = max(r, tmp);
		}
		cout << r << endl;
	}
	return 0;
}
