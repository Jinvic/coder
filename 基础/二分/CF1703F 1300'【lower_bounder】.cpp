// https://codeforces.com/contest/1703/problem/F
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
const int N = 200005;
const int inf = 1000000009;
int a[N];
vector<int> v;
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);
	int t, n, x;
	cin >> t;
	while (t--)
	{
		cin >> n;
		for (int i = 1; i <= n; i++)
			cin >> a[i];
		ll res = 0;
		v.clear();
		for (int i = 1; i <= n; i++)
		{
			if (a[i] < i)
			{
				res += lower_bound(v.begin(), v.end(), a[i]) - v.begin();
				v.push_back(i);
			}
		}
		cout << res << endl;
	}
	return 0;
}
