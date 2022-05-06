//https://codeforces.com/problemset/problem/180/C
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 100005;
const int inf = 0x7f7f7f7f;
const int mod = 1000000007;
const double eps = 1e-8;
int dp[2][2];
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	string s;
	cin >> s;
	s = " " + s;
	int p = 1;
	for (int i = 1; i <= s.length(); i++)
	{
		if (isupper(s[i]))
		{
			dp[p][1] = dp[!p][1];					  //第i位为大写
			dp[p][0] = min(dp[!p][0], dp[!p][1]) + 1; //第i位为小写
		}
		else
		{
			dp[p][1] = dp[!p][1] + 1;			  //第i位为大写
			dp[p][0] = min(dp[!p][0], dp[!p][1]); //第i位为小写
		}
		p = !p;
	}
	cout << min(dp[!p][0], dp[!p][1]);
	return 0;
}
