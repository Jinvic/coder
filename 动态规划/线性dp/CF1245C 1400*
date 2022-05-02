//https://codeforces.com/problemset/problem/1245/C
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 100005;
const int inf = 0x7f7f7f7f;
const int mod = 1000000007;
const double eps = 1e-8;
int dp[N];
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);

	string s;
	cin >> s;
	dp[0] = dp[1] = 1;
	s = " " + s;
	if (s[1] == 'm' || s[1] == 'w')
	{
		cout << 0 << endl;
		return 0;
	}
	for (int i = 2; i < s.size(); i++)
	{
		if (s[i] == 'm' || s[i] == 'w')
		{
			cout << 0 << endl;
			return 0;
		}
		if (s[i] == s[i - 1] && (s[i] == 'u' || s[i] == 'n'))
			dp[i] = (dp[i - 1] + dp[i - 2]) % mod;
		else
			dp[i] = dp[i - 1];
	}
	cout << dp[s.size() - 1] << endl;
	return 0;
}
