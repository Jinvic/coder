//https://codeforces.com/contest/1703/problem/G
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
const int N = 100005;
const int inf = 1000000009;
ll a[N];
const int GAP = 30; // log(10^9)约30
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	int t, n, k;
	cin >> t;
	while (t--)
	{
		cin >> n >> k;
		for (int i = 0; i < n; i++)
			cin >> a[i];
		ll sum = 0;
		ll res = 0;
		ll tmp = 0;
		ll cost = 0;
		int len = min(n, GAP);
		for (int i = 0; i < len; i++)
			tmp += a[i] >> (i + 1);
		res = max(res, tmp);
		for (int i = 0; i < n; i++)
		{
			sum += a[i];
			cost += k;
			tmp = 0;
			len = min(n - i - 1, GAP);
			for (int j = 0; j < len; j++)
				tmp += a[i + j + 1] >> (j + 1);
			res = max(res, sum - cost + tmp);
		}
		cout << res << endl;
	}
	return 0;
}
