//https://codeforces.com/contest/698/problem/A
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 100005;
const int inf = 0x7f7f7f7f;
const int mod = 32768;
const double eps = 1e-8;

int dp[105][3]; //第i天 做什么 最大活动天数
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	int n, x;
	cin >> n;
	bool p1 = false, p2 = false;
	for (int i = 1; i <= n; i++)
	{
		cin >> x;
		if (!x)
		{
			dp[i][1] = dp[i][2] = max(dp[i - 1][1], dp[i - 1][2]);
			p2 = p1 = false;
		}
		else
		{
			bool f1 = (x == 1 || x == 3), f2 = (x == 2 || x == 3);
			dp[i][1] = dp[i - 1][2] + f1;
			dp[i][2] = dp[i - 1][1] + f2;
			if (!p1) //前一天没做1，今天也可以做1
				dp[i][1] = max(dp[i][1], dp[i - 1][1] + f1);
			if (!p2) //同上
				dp[i][2] = max(dp[i][2], dp[i - 1][2] + f2);
			p1 = f1, p2 = f2;
		}
	}
	cout << n - max(dp[n][1], dp[n][2]);
	return 0;
}
