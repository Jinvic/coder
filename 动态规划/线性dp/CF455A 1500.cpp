//https://codeforces.com/problemset/problem/455/A
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 100005;
const int inf = 0x7f7f7f7f;
const int mod = 32768;
const double eps = 1e-8;

ll a[N];
ll dp[N];
map<ll, int> cnt;
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	int n, m, x;
	cin >> n;
	int len = 0;
	for (int i = 0; i < n; i++)
	{

		cin >> x;
		if (!cnt[x])
		{
			cnt[x] = 1;
			a[len++] = x;
		}
		else
			cnt[x]++;
	}
	sort(a, a + len);
	if (len == 1)
	{
		cout << a[0] * n << endl;
		return 0;
	}
	else if (len == 2)
	{
		if (a[0] != a[1] - 1) //两个都取
			cout << a[0] * cnt[a[0]] + a[1] * cnt[a[1]];
		else //取更大值
			cout << max(a[0] * cnt[a[0]], a[1] * cnt[a[1]]);
		return 0;
	} //注意，对每相邻两个取更大值
	dp[0] = a[0] * cnt[a[0]];
	dp[1] = a[1] * cnt[a[1]];
	dp[2] = a[2] * cnt[a[2]];
	if (a[0] != a[1] - 1)
		dp[1] += dp[0];
	if (a[1] != a[2] - 1) //取更大值
		dp[2] += max(dp[1], dp[0]);
	else //只能取dp[0]
		dp[2] += dp[0];
	for (int i = 3; i < len; i++)
	{
		if (a[i] == a[i - 1] + 1)
			dp[i] = a[i] * cnt[a[i]] + max(dp[i - 2], dp[i - 3]); //不能取dp[i-1]，取更前两位的更大值
		else
			dp[i] = a[i] * cnt[a[i]] + max(dp[i - 1], dp[i - 2]);
	}
	cout << max(dp[len - 1], dp[len - 2]) << endl;
	return 0;
}
