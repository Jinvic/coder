//https://codeforces.com/contest/1673/problem/C
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 40004;
const int inf = 0x7f7f7f7f;
const int mod = 1000000007;
const double eps = 1e-8;

int dp[N][505];
int a[N], len;
int digit[6];
bool judge(int x)
{
	int cnt = 0;
	while (x)
	{
		digit[cnt++] = x % 10;
		x /= 10;
	}
	int i = 0, j = cnt - 1;
	while (i < j)
	{
		if (digit[i] != digit[j])
			return false;
		i++, j--;
	}
	return true;
}
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	len = 0;
	for (int i = 1; i <= N; i++)
	{
		if (judge(i))
		{
			// cout << i << endl;
			a[len++] = i;
		}
	}
	// cout << len << " " << a[len - 1];
	for (int i = 0; i < len; i++)
		dp[0][i] = 1;
	for (int i = 1; i < N; i++)
	{
		for (int j = 0; j < len; j++)
			if (a[j] <= i)
				dp[i][j] = (dp[i][j - 1] + dp[i - a[j]][j]) % mod;
			else
				dp[i][j] = dp[i][j - 1];
	}
	int t;
	int n;
	cin >> t;
	while (t--)
	{
		cin >> n;
		cout << dp[n][len - 1] << endl;
	}
	return 0;
}
