//https://ac.nowcoder.com/acm/problem/235687
#include <bits/stdc++.h>
using namespace std;
typedef long long ll;
typedef pair<int, int> pii;
const int N = 200005;
const ll mod = 1000000007;
const double eps = 1e-8;
ll a[N];
vector<ll> p, m;
signed main(void)
{
	ios::sync_with_stdio(false);
	cin.tie(0), cout.tie(0);

	int n;
	cin >> n;
	for (int i = 0; i < n; i++)
		cin >> a[i];
	sort(a, a + n);
	for (int i = 0; i < n / 2; i++)
		p.push_back(a[i]);
	for (int i = n / 2; i < n; i++)
		m.push_back(a[i]);
	reverse(p.begin(), p.end());
	ll r = 1;
	int i = 0, j = 0;
	while (i < m.size() || j < p.size())
	{
		if (i < m.size())
			r = (r * m[i]) % mod;
		if (j < p.size())
			r = (r + p[i]) % mod;
		i++, j++;
	}
	cout << r << endl;
	return 0;
}
