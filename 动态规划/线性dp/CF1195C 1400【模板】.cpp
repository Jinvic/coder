//https://codeforces.com/contest/1195/problem/C
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 100005;
const int inf = 0x7f7f7f7f;
const int mod = 32768;
const double eps = 1e-8;

ll dp[2][3]; //滚动数组优化
ll h1[N], h2[N];
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	int n;
	cin >> n;
	for (int i = 1; i <= n; i++)
		cin >> h1[i];
	for (int i = 1; i <= n; i++)
		cin >> h2[i];
	int p = 1;
	for (int i = 1; i <= n; i++)
	{
		dp[p][0] = max(dp[!p][0], max(dp[!p][1], dp[!p][2])); //这次不拿，之前所有情况都可选
		dp[p][1] = max(dp[!p][0], dp[!p][2]) + h1[i];		  //之前没拿或拿2
		dp[p][2] = max(dp[!p][0], dp[!p][1]) + h2[i];		  //之前没拿或拿1
		p = !p;
	}
	cout << max(dp[!p][0], max(dp[!p][1], dp[!p][2])) << endl;
	return 0;
}
